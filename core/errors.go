package core

import "github.com/pkg/errors"

var (
	ErrAlphaVantageKeyNotDefined = errors.New("It was impossible to load AlphaVantageKey env var")
	ErrStocksNATSKeyNotDefined   = errors.New("It was impossible to load Stocks NATS URI env var")
)
