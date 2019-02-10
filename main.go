package main

import (
	"fmt"
	"github.com/pedrolopesme/sentinel/core"
	"go.uber.org/zap"
	"os"
)

var (
	sentinelConfig *core.SentinelConfig
	logger         *zap.Logger
	err            error
)

func init() {
	// Loading sentinel configs
	sentinelConfig, err = core.NewSentinelConfig()
	if err != nil {
		fmt.Println("Fail to load sentinel config.")
		os.Exit(1)
	}

	logger, err = initializeLogger(sentinelConfig)
	if err != nil {
		fmt.Printf("It was impossible to load logger. Killing sentinel. Error: %v", err.Error())
		os.Exit(1)
	}
}

// final runs when the Sentinel terminates
func final() {
	defer func() {
		if err := logger.Sync(); err != nil {
			fmt.Println("It was impossible to flush the logger")
			os.Exit(1)
		}
	}()
}

func initializeLogger(config *core.SentinelConfig) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		fmt.Sprintf("%v/sentinels.log", config.LogsPath),
	}
	return cfg.Build()
}

func main() {
	defer final()
	printLogo()

	// Hardcoding a stock to test sentinel
	// TODO: replace this with something more flexible.
	var (
		schedule = core.NewSchedule("PETR3.SA", "1min")
		sentinel = core.NewStockSentinel(sentinelConfig, schedule)
	)

	// Running sentinel
	executionId, err := sentinel.Run()
	if err != nil {
		logger.Error("Fail to run sentinel",
			zap.String("sentinelId", sentinel.GetId()),
			zap.String("executionId", executionId),
			zap.String("method", "main"),
			zap.String("error", err.Error()))
		os.Exit(1)
	}

	logger.Info("Sentinel have run successfully",
		zap.String("sentinelId", sentinel.GetId()),
		zap.String("executionId", executionId),
		zap.String("method", "main"))

	// Trying to kill sentinel
	if err := sentinel.Kill(); err != nil {
		logger.Error("Fail to kill sentinel",
			zap.String("sentinelId", sentinel.GetId()),
			zap.String("executionId", executionId),
			zap.String("error", err.Error()))
		os.Exit(1)
	}

	logger.Info("Sentinel have terminated successfully",
		zap.String("sentinelId", sentinel.GetId()),
		zap.String("executionId", executionId),
		zap.String("method", "main"))
}

func printLogo() {
	fmt.Println(`
	________  _______  _____  ___  ___________  __    _____  ___    _______  ___       
 /"       )/"     "|(\"   \|"  \("     _   ")|" \  (\"   \|"  \  /"     "||"  |      
(:   \___/(: ______)|.\\   \    |)__/  \\__/ ||  | |.\\   \    |(: ______)||  |      
 \___  \   \/    |  |: \.   \\  |   \\_ /    |:  | |: \.   \\  | \/    |  |:  |      
  __/  \\  // ___)_ |.  \    \. |   |.  |    |.  | |.  \    \. | // ___)_  \  |___   
 /" \   :)(:      "||    \    \ |   \:  |    /\  |\|    \    \ |(:      "|( \_|:  \  
(_______/  \_______) \___|\____\)    \__|   (__\_|_)\___|\____\) \_______) \_______) 
              
=========================== 
       Up and Running
===========================`)
}
