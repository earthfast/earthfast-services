package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"sync"
	"syscall"
	"time"

	"armada-node/api"
	contentapi "armada-node/api/content"
	"armada-node/api/middleware"
	"armada-node/hosting"
	"armada-node/hosting/site"
	"armada-node/logging"
	logshttp "armada-node/logging/http"
	"armada-node/metering/uptime"
	"armada-node/model"
	"armada-node/model/cache"
	"armada-node/model/dynamic"

	"contrib.go.opencensus.io/exporter/prometheus"
	"go.etcd.io/bbolt"
	"go.opencensus.io/stats/view"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	flags = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	contractAddress = flags.String(
		"contract-address",
		os.Getenv("CONTRACT_ADDRESS"),
		"Hex address where the smart contract is deployed",
	)

	databaseDir = flags.String(
		"database-dir",
		os.Getenv("DATABASE_DIR"),
		"The directory where databases are stored, must persist across reboots",
	)

	ethRPCEndpoint = flags.String(
		"eth-rpc-endpoint",
		os.Getenv("ETH_RPC_ENDPOINT"),
		"URL at which to make Ethereum RPCs",
	)

	hostingCacheDir = flags.String(
		"hosting-cache-dir",
		os.Getenv("HOSTING_CACHE_DIR"),
		"The directory where static site content is stored",
	)

	httpPort = flags.Int(
		"http-port",
		mustAtoi(os.Getenv("HTTP_PORT")),
		"Port to bind the HTTP server",
	)

	logLevel = flags.String(
		"log-level",
		os.Getenv("LOG_LEVEL"),
		"The minimum enabled logging level, case-insensitive (e.g. debug, info, error)",
	)

	nodeIDStr = flags.String(
		"node-id",
		os.Getenv("NODE_ID"),
		"The ID of this node, in hexadecimal",
	)
)

var (
	// gracefulTimeout is the maximum duration of a graceful shutdown.
	gracefulTimeout = 5 * time.Second
)

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.Encoding = "console"
	loggerConfig.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format("2006-01-02T15:04:05Z") + " [UTC]")
	}

	loggerConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	logger, err := logging.CreateCaptureLogger(loggerConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}

	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	// Parse flags
	flags.Parse(os.Args[1:])

	// Log flags using Zap
	logger.Info("Starting content node with configuration:")
	flags.VisitAll(func(f *flag.Flag) {
		logger.Info(fmt.Sprintf("%s=%s", f.Name, f.Value))
	})

	// Update log level after flags are parsed
	if *logLevel != "" {
		if lvl, err := zap.ParseAtomicLevel(*logLevel); err == nil {
			// Rebuild logger with development config for debug level
			if lvl.Level() == zapcore.DebugLevel {
				loggerConfig = zap.NewDevelopmentConfig()
				loggerConfig.Encoding = "console"
				loggerConfig.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
					enc.AppendString(t.UTC().Format("2006-01-02T15:04:05Z") + " [UTC]")
				}
				loggerConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
			}

			loggerConfig.Level = lvl
			updatedLogger, err := logging.CreateCaptureLogger(loggerConfig)
			if err != nil {
				logger.Error("Failed to update log level", zap.Error(err))
			} else {
				logger.Info("Setting log level", zap.String("level", lvl.String()))
				zap.ReplaceGlobals(updatedLogger)
				logger = updatedLogger
			}
		} else {
			logger.Error("Invalid log level", zap.String("level", *logLevel), zap.Error(err))
		}
	}

	// Validate flags.
	nodeID, err := model.ParseID(*nodeIDStr)
	if err != nil {
		logger.Fatal("Invalid node-id", zap.Error(err))
	}

	// Ensure that all directories specified as flags exist, creating them if necessary.
	if err := os.MkdirAll(*databaseDir, 0755); err != nil && !os.IsExist(err) {
		logger.Fatal("Failed to initialize database directory", zap.Error(err))
	}
	if err := os.MkdirAll(*hostingCacheDir, 0755); err != nil && !os.IsExist(err) {
		logger.Fatal("Failed to initialize hosting cache directory", zap.Error(err))
	}

	// Initialize metrics
	if err := api.RegisterMetrics(); err != nil {
		logger.Fatal("Failed to register API metrics", zap.Error(err))
	}

	// Create and register Prometheus exporter
	pe, err := prometheus.NewExporter(prometheus.Options{
		Namespace: "earthfast",
		OnError: func(err error) {
			logger.Error("Prometheus export error", zap.Error(err))
		},
	})
	if err != nil {
		logger.Fatal("Failed to create Prometheus exporter", zap.Error(err))
	}
	view.RegisterExporter(pe)

	// Set metrics reporting period
	if *logLevel == "debug" {
		view.SetReportingPeriod(1 * time.Second)
		logger.Debug("Metrics reporting enabled with 1s interval")
	} else {
		view.SetReportingPeriod(10 * time.Second)
	}

	// Initialize the model
	var modelClient model.Client
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	if modelClient, err = model.NewEthClient(
		ctx,
		model.EthClientArgs{
			Endpoint: *ethRPCEndpoint,
			Address:  *contractAddress,
			Logger:   logger,
		},
		model.EthClientOptions{},
	); err != nil {
		logger.Fatal("Failed to initialize Ethereum client", zap.Error(err))
	}
	cancel()

	// Cache the model in local memory
	modelClient = cache.NewClient(modelClient, cache.Options{})

	// Setup a dynamic.Node that holds this content node's latest on-chain state.
	// It will periodically refresh itself in order to pick up changes, such as being
	// assigned to a new project.
	//
	// Bootstrapping the dynamic.Node blocks server startup because without at least
	// one successful fetch we can't know which project to serve (if any).
	self := dynamic.NewNode(logger, modelClient, nodeID, dynamic.Options{})
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	if err := self.Bootstrap(ctx); err != nil {
		logger.Fatal("Failed to bootstrap dynamic node", zap.Error(err))
	}
	cancel()
	self.Start()
	defer self.Stop()

	// Uptime Meter
	uptimeStore, err := uptime.NewBoltDBStore(
		logger,
		filepath.Join(*databaseDir, "uptime.db"),
		&bbolt.Options{Timeout: 15 * time.Second},
	)
	if err != nil {
		logger.Fatal("Failed to initialize uptime store", zap.Error(err))
	}
	defer uptimeStore.Close()
	uptimeMeter, err := uptime.NewMeter(
		uptime.MeterArgs{
			Model:  modelClient,
			Self:   self,
			Store:  uptimeStore,
			Logger: logger,
		},
		uptime.MeterOptions{},
	)
	if err != nil {
		logger.Fatal("Failed to initialize uptime meter", zap.Error(err))
	}
	uptimeMeter.Start()

	// Initialize the ProjectManager to handle automatic startup of Next.js projects
	projectManager := hosting.NewProjectManager(
		logger,
		modelClient,
		*hostingCacheDir,
		hosting.DefaultSiteProvider(logger, modelClient, *hostingCacheDir),
	)

	// Initialize existing projects assigned to this node
	if node := self.Get(); node != nil {
		logger.Info("Checking for assigned projects to initialize")
		if err := projectManager.InitializeFromNode(context.Background(), node); err != nil {
			logger.Error("Failed to initialize projects from node state", zap.Error(err))
		}
	}

	// Add listener to initialize projects when node state changes (new project assignment)
	self.AddListener(func(node *model.Node) {
		if node != nil && !node.ProjectID.IsZero() {
			logger.Info("Node state changed, checking project assignment",
				zap.String("projectId", node.ProjectID.Hex()))
			if err := projectManager.InitializeProject(context.Background(), node.ProjectID); err != nil {
				logger.Error("Failed to initialize project after node update", zap.Error(err))
			}
		}
	})

	// Ensure ProjectManager is stopped during shutdown
	defer projectManager.Stop()
	defer site.StopAllNextJSProcesses()

	// API Handler
	apiHandler := contentapi.NewHandler(
		logger,
		modelClient,
		uptimeMeter,
		site.NewServeMux(
			logger,
			modelClient,
			hosting.DefaultSiteProvider(logger, modelClient, *hostingCacheDir),
		),
		self,
	)

	// Debug metrics handler
	debugMetricsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := view.RetrieveData("service/content/total_requests")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Metrics data: %+v\n", data)
	})

	// Handler for /statusz
	statuszHandler := contentapi.NewStatuszHandler(logger, self)

	// HTTP Server
	mux := http.NewServeMux()
	mux.Handle("/", apiHandler)
	mux.Handle("/metrics", pe)
	mux.Handle("/debug/metrics", debugMetricsHandler)

	mux.Handle("/statusz", statuszHandler)

	mux.Handle("/api/logs", logshttp.LogsHandler(logger, self))
	mux.Handle("/logs", logshttp.LogsViewerHandler(self))

	middlewareChain := middleware.Chain(
		middleware.WithLogger(logger),
		middleware.WithOpenCensus(),
		middleware.WithCORS(),
	)

	// Apply the middleware chain to your mux
	handler := middlewareChain(mux)

	httpSrv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", *httpPort),
		Handler: handler,
	}

	// Start the HTTP server
	var wg sync.WaitGroup
	stopCh := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()

		logger.Info("Serving HTTP", zap.String("address", httpSrv.Addr))
		if err := httpSrv.ListenAndServe(); err != nil {
			select {
			case <-stopCh:
				// Expected shutdown, no need to log as error
			default:
				logger.Fatal("HTTP server error", zap.Error(err))
			}
		}
	}()

	// Wait for shutdown signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	logger.Info("Caught SIGTERM, shutting down...")
	close(stopCh)

	// Shutdown the HTTP server.
	// Allow the server some time to finish serving inflight requests, after
	// which any remaining active connections will be forcefully closed.
	wg.Add(1)
	go func() {
		defer wg.Done()

		ctx, cancel := context.WithTimeout(context.Background(), gracefulTimeout)
		defer cancel()
		if err := httpSrv.Shutdown(ctx); err != nil {
			logger.Warn("HTTP server forced shutdown", zap.Error(err))
			httpSrv.Close()
		}
	}()
	wg.Wait()
	logger.Info("Shutdown complete")
}
