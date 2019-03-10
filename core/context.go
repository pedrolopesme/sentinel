package core

import (
	"fmt"
	"github.com/pedrolopesme/sentinel/client"
	"github.com/satori/go.uuid"
)

type Context interface {
	GetSentinelConfig() *SentinelConfig
	GetStockNats() client.NATSServer
}

// AppContext represents a general context for the application
type AppContext struct {
	sentinelConfig  *SentinelConfig
	stockNatsServer client.NATSServer
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

	return &AppContext{
		stockNatsServer: stockNATS,
		sentinelConfig:  config,
	}, nil
}

// GetSentinelConfig returns Sentinel config
func (ac *AppContext) GetSentinelConfig() *SentinelConfig {
	return ac.sentinelConfig
}

// GetStockNats returns Nats server connection
func (ac *AppContext) GetStockNats() client.NATSServer {
	return ac.stockNatsServer
}
