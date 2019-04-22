package core

import (
	"fmt"
	"github.com/pedrolopesme/sentinel/client"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"os"
)

type Context interface {
	SentinelConfig() *SentinelConfig
	StockNats() client.NATSServer
	Logger() *zap.Logger
}

// AppContext represents a general context for the application
type AppContext struct {
	sentinelConfig  *SentinelConfig
	stockNatsServer client.NATSServer
	logger          *zap.Logger
}

// SentinelConfig returns Sentinel config
func (ac *AppContext) SentinelConfig() *SentinelConfig {
	return ac.sentinelConfig
}

// StockNats returns Nats server connection
func (ac *AppContext) StockNats() client.NATSServer {
	return ac.stockNatsServer
}

// Logger returns application default logger
func (ac *AppContext) Logger() *zap.Logger {
	return ac.logger
}

func initializeLogger(config *SentinelConfig) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		fmt.Sprintf("%v/sentinels.log", config.LogsPath),
	}
	return cfg.Build()
}

// NewAppContext knows how to instantiate Sentinels General Context
// TODO add some logging
// TODO add tests
func NewAppContext(config *SentinelConfig) (ctx *AppContext, err error) {
	clientID, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Cant get stocks due to", err.Error())
	}

	stockNATS, err := client.NewNATSServer(config.NATSStocksClusterID, clientID.String(), config.NATSStocksURI)
	if err != nil {
		return nil, err
	}

	logger, err := initializeLogger(config)
	if err != nil {
		fmt.Printf("It was impossible to load logger. Killing sentinel. Error: %v", err.Error())
		os.Exit(1)
	}

	return &AppContext{
		stockNatsServer: stockNATS,
		sentinelConfig:  config,
		logger:          logger,
	}, nil
}
