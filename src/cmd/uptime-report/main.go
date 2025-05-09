package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"armada-node/model"
	"armada-node/reporting/uptime"

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

	endTime = flags.Int64(
		"end-time",
		0,
		"The report end time",
	)

	ethRPCEndpoint = flags.String(
		"eth-rpc-endpoint",
		os.Getenv("ETH_RPC_ENDPOINT"),
		"URL at which to make Ethereum RPCs",
	)

	nominalUptimeRatioPercentile = flags.Int(
		"nominal-uptime-ratio-percentile",
		-1,
		"See ReportOptions.NominalUptimeRatioPercentile",
	)

	probeFailureCountPercentile = flags.Int(
		"probe-failure-count-percentile",
		-1,
		"See ReportOptions.ProbeFailureCountPercentile",
	)

	projectIDStr = flags.String(
		"project-id",
		"",
		"Hex representation of the project ID to generate a report about",
	)

	requestCountPercentile = flags.Int(
		"request-count-percentile",
		-1,
		"See ReportOptions.RequestCountPercentile",
	)

	startTime = flags.Int64(
		"start-time",
		0,
		"The report start time",
	)

	reportTimeout = flags.Duration(
		"report-timeout",
		time.Minute,
		"The maximum amount of time allowed to generate the report",
	)

	logLevel = flags.String(
		"log-level",
		os.Getenv("LOG_LEVEL"),
		"The minimum enabled logging level, case-insensitive (e.g. debug, info, error)",
	)
)

func main() {
	loggerConfig := zap.NewDevelopmentConfig()
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
	logger.Info("Starting uptime reporter with configuration:")
	flags.VisitAll(func(f *flag.Flag) {
		logger.Info(fmt.Sprintf("%s=%s", f.Name, f.Value))
	})

	// Update log level after flags are parsed
	if *logLevel != "" {
		if lvl, err := zap.ParseAtomicLevel(*logLevel); err == nil {
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

	// Initialize the model
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	modelClient, err := model.NewEthClient(
		ctx,
		model.EthClientArgs{
			Endpoint: *ethRPCEndpoint,
			Address:  *contractAddress,
			Logger:   logger,
		},
		model.EthClientOptions{},
	)
	if err != nil {
		logger.Fatal("Failed to initialize Ethereum client", zap.Error(err))
	}
	cancel()

	// Parse the project ID
	projectID, err := model.ParseID(*projectIDStr)
	if err != nil {
		logger.Fatal("Failed to parse project ID",
			zap.String("projectId", *projectIDStr),
			zap.Error(err))
	}

	// Create the reporter
	reportOptions := uptime.ReportOptions{
		StartTime:                    time.Unix(*startTime, 0),
		EndTime:                      time.Unix(*endTime, 0),
		RequestCountPercentile:       *requestCountPercentile,
		ProbeFailureCountPercentile:  *probeFailureCountPercentile,
		NominalUptimeRatioPercentile: *nominalUptimeRatioPercentile,
	}

	logger.Info("Creating uptime reporter",
		zap.Time("startTime", reportOptions.StartTime),
		zap.Time("endTime", reportOptions.EndTime),
		zap.Int("requestCountPercentile", reportOptions.RequestCountPercentile),
		zap.Int("probeFailureCountPercentile", reportOptions.ProbeFailureCountPercentile),
		zap.Int("nominalUptimeRatioPercentile", reportOptions.NominalUptimeRatioPercentile))

	reporter, err := uptime.NewReporter(
		uptime.NewContentNodeProvider(modelClient, &http.Client{Timeout: time.Minute}),
		reportOptions,
	)
	if err != nil {
		logger.Fatal("Failed to create uptime reporter", zap.Error(err))
	}

	// Run the report
	logger.Info("Running uptime report",
		zap.String("projectId", projectID.Hex()),
		zap.Duration("timeout", *reportTimeout))

	ctx, cancel = context.WithTimeout(context.Background(), *reportTimeout)
	defer cancel()
	result, err := reporter.Run(ctx, projectID)
	if err != nil {
		logger.Fatal("Failed to run report", zap.Error(err))
	}

	// Print the report to stdout
	logger.Info("Report generated successfully, printing to stdout")
	prettyJSON := json.NewEncoder(os.Stdout)
	prettyJSON.SetIndent("", "    ")
	if err := prettyJSON.Encode(&result); err != nil {
		logger.Error("Error printing report", zap.Error(err))
	}

	logger.Info("Report completed successfully")
}
