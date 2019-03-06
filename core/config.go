package core

import (
	"os"
)

const (
	ALPHA_VANTAGE_ENV_VAR_NAME     = "ALPHAVANTAGE_KEY"
	LOGS_PATH_ENV_VAR_NAME         = "SENTINEL_LOGS_PATH"
	NATS_STOCKS_VAR_NAME           = "NATS_STOCKS_DATA_URI"
	NATS_STOCKS_CLUSTERID_VAR_NAME = "NATS_STOCKS_CLUSTERID"
	DEFAULT_LOGS_PATH              = "./logs"
)

// SentinelConfig stores the general configuration directives for a Sentinel to run
type SentinelConfig struct {
	AlphaVantageKey     string
	LogsPath            string
	NATSStocksURI       string
	NATSStocksClusterID string
}

// NewSentinelConfig creates an instance of Sentinel Config, loading
// all data necessary to run a Sentinel
// TODO add tests
func NewSentinelConfig() (config *SentinelConfig, err error) {
	alphaVantageKey, err := getAlphaVantageKey()
	if err != nil {
		return nil, err
	}

	stocksNATSURI, err := getStocksNATSURI()
	if err != nil {
		return nil, err
	}

	stocksNATSClusterID, err := getStocksNATSClusterId()
	if err != nil {
		return nil, err
	}

	config = &SentinelConfig{
		LogsPath:            getLogsPath(),
		AlphaVantageKey:     alphaVantageKey,
		NATSStocksURI:       stocksNATSURI,
		NATSStocksClusterID: stocksNATSClusterID,
	}
	return
}

// getLogsPath tries to identify where to store log files. First,
// it tires to get that info from Env Vars, then it assumes a default configuration
// TODO add tests
func getLogsPath() (logsPath string) {
	logsPath = os.Getenv(LOGS_PATH_ENV_VAR_NAME)
	if logsPath == "" {
		logsPath = DEFAULT_LOGS_PATH
	}
	return
}

// getAlphaVantageKey tries to load AlphaVantage Key env var
// TODO add tests
func getAlphaVantageKey() (key string, err error) {
	key = os.Getenv(ALPHA_VANTAGE_ENV_VAR_NAME)
	if key == "" {
		return "", ErrAlphaVantageKeyNotDefined
	}
	return
}

// getStocksNATSURI tries to load NATS Server URI.
// This server is responsible for storing stocks data collected by a Sentinel
// TODO add tests
func getStocksNATSURI() (uri string, err error) {
	uri = os.Getenv(NATS_STOCKS_VAR_NAME)
	if uri == "" {
		return "", ErrStocksNATSKeyNotDefined
	}
	return
}

// getStocksNATSClusterId tries to load NATS Server Cluster ID.
// TODO add tests
func getStocksNATSClusterId() (clusterID string, err error) {
	clusterID = os.Getenv(NATS_STOCKS_CLUSTERID_VAR_NAME)
	if clusterID == "" {
		return "", ErrStocksNATSClusterIDKeyNotDefined
	}
	return
}
