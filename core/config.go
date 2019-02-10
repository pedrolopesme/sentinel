package core

import (
	"os"
)

const (
	ALPHA_VANTAGE_ENV_VAR_NAME = "ALPHAVANTAGE_KEY"
	LOGS_PATH_ENV_VAR_NAME     = "SENTINEL_LOGS_PATH"
	DEFAULT_LOGS_PATH          = "./logs"
)

// SentinelConfig stores the general configuration directives for a Sentinel to run
type SentinelConfig struct {
	AlphaVantageKey string
	LogsPath        string
}

func NewSentinelConfig() (config *SentinelConfig, err error) {
	alphaVantageKey, err := getAlphaVantageKey()
	if err != nil {
		return nil, err
	}
	config = &SentinelConfig{
		LogsPath:        getLogsPath(),
		AlphaVantageKey: alphaVantageKey,
	}

	return
}

// getLogsPath tries to identify where to store log files. First,
// it tires to get that info from Env Vars, then it assumes a default configuration
func getLogsPath() (logsPath string) {
	logsPath = os.Getenv(LOGS_PATH_ENV_VAR_NAME)
	if logsPath == "" {
		logsPath = DEFAULT_LOGS_PATH
	}
	return
}

// getLogsPath tries to identify where to store log files. First,
// it tires to get that info from Env Vars, then it assumes a default configuration
func getAlphaVantageKey() (key string, err error) {
	key = os.Getenv(ALPHA_VANTAGE_ENV_VAR_NAME)
	if key == "" {
		return "", ErrAlphaVantageKeyNotDefined
	}
	return
}
