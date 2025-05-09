package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"text/template"
	"time"

	domainapi "armada-node/api/domain"
	"armada-node/api/middleware"
	"armada-node/geo"
	"armada-node/model"
	"armada-node/model/cache"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	flags = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	contractAddress = flags.String(
		"contract-address",
		os.Getenv("CONTRACT_ADDRESS"),
		"Hex address where the registry smart contract is deployed",
	)

	domainToProjectMapping = flags.String(
		"domain-to-project-mapping",
		os.Getenv("DOMAIN_TO_PROJECT_MAPPING"),
		"Map where the key is a domain name (or '*') and the value is a project ID. Usage: '--domain-to-project-mapping app.example.com=0x123...,foo.com=0x456...'",
	)

	domainToProjectMappingUrl = flags.String(
		"domain-to-project-mapping-url",
		os.Getenv("DOMAIN_TO_PROJECT_MAPPING_URL"),
		"URL to a JSON file containing domain-to-project mappings. The file must contain an array of objects, each with 'url' and 'projectId' fields. Example: [{'url':'app.example.com','projectId':'0x123...'},{'url':'foo.com','projectId':'0x456...'}]'",
	)

	ethRPCEndpoint = flags.String(
		"eth-rpc-endpoint",
		os.Getenv("ETH_RPC_ENDPOINT"),
		"URL at which to make Ethereum RPCs",
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

	webTemplateDir = flags.String(
		"web-templates-dir",
		os.Getenv("WEB_TEMPLATES_DIR"),
		"The directory where templates for web content are stored",
	)

	webStaticDir = flags.String(
		"web-static-dir",
		os.Getenv("WEB_STATIC_DIR"),
		"The directory where static web content is stored",
	)

	ipLookupAPIKey = flags.String(
		"ip-lookup-api-key",
		os.Getenv("IP_LOOKUP_API_KEY"),
		"The API key for the IP lookup service",
	)

	environment = flags.String(
		"environment",
		func() string {
			env := os.Getenv("ENVIRONMENT")
			if env != "" {
				return env
			}
			return "production"
		}(),
		"The environment in which the domain node is running",
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
	logger, err := loggerConfig.Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}

	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	// Parse flags
	flags.Parse(os.Args[1:])

	// Log flags using Zap
	logger.Info("Starting domain node with configuration:")
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
			updatedLogger, err := loggerConfig.Build()
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

	ctx := context.Background()

	// Initialize the model
	var modelClient model.Client
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
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

	// Domain-to-Project resolver
	var resolver domainapi.Resolver
	logger.Debug("Selecting domain-to-project mapping method",
		zap.String("mapping", *domainToProjectMapping),
		zap.String("mappingUrl", *domainToProjectMappingUrl))

	if *domainToProjectMappingUrl != "" {
		logger.Info("Downloading and parsing domain-to-project-mapping from URL",
			zap.String("url", *domainToProjectMappingUrl))
		resolver, err = parseDomainToProjectMappingUrl(*domainToProjectMappingUrl, logger)
		if err != nil {
			logger.Fatal("Failed to parse domain-to-project-mapping-url", zap.Error(err))
		}
	} else if *domainToProjectMapping != "" {
		logger.Info("Parsing domain-to-project-mapping from configuration")
		resolver, err = parseDomainToProjectMapping()
		if err != nil {
			logger.Fatal("Failed to parse domain-to-project-mapping", zap.Error(err))
		}
	} else {
		logger.Fatal("Either --domain-to-project-mapping or --domain-to-project-mapping-url must be provided")
	}

	// Prepare templates that will be served by the API handler
	swTemplate, err := template.ParseFiles(filepath.Join(*webTemplateDir, "main.js.tmpl"))
	if err != nil {
		logger.Fatal("Failed to parse service worker template",
			zap.String("path", filepath.Join(*webTemplateDir, "main.js.tmpl")),
			zap.Error(err))
	}

	abstractClient, err := geo.NewAbstractClient(*ipLookupAPIKey, logger)
	if err != nil {
		logger.Warn("Failed to initialize Abstract IP Geolocation Client", zap.Error(err))
	}

	apiHandler, err := domainapi.NewHandler(
		logger,
		modelClient,
		resolver,
		domainapi.Templates{ServiceWorker: swTemplate},
		http.Dir(*webStaticDir),
		abstractClient,
		*environment,
	)
	if err != nil {
		logger.Fatal("Failed to initialize API handler", zap.Error(err))
	}

	middlewareChain := middleware.Chain(
		middleware.WithLogger(logger),
		middleware.WithRealIP(),
		middleware.WithCORS(),
	)

	httpSrv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", *httpPort),
		Handler: middlewareChain(apiHandler),
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

func parseDomainToProjectMapping() (*domainapi.StaticResolver, error) {
	input := strings.TrimSpace(*domainToProjectMapping)
	if input == "" {
		return nil, errors.New("at least one entry is required")
	}

	data := make(map[string]model.ID)
	for _, entry := range strings.Split(input, ",") {
		parts := strings.Split(entry, "=")
		if len(parts) != 2 {
			return nil, fmt.Errorf("malformed entry: %s", entry)
		}
		domain := strings.TrimSpace(parts[0])
		projectID, err := model.ParseID(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, fmt.Errorf("parsing project ID: %v", err)
		}
		data[domain] = projectID
	}
	return domainapi.NewStaticResolver(data), nil
}

func parseDomainToProjectMappingUrl(url string, logger *zap.Logger) (*domainapi.DynamicResolver, error) {
	return domainapi.NewDynamicResolver(url, logger), nil
}
