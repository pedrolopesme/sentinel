package client

import "github.com/pkg/errors"

var (
	ErrCantGetStockPricesFromAlphaVantage = errors.New("It was impossible to retrieve stock prices from Alpha Vantage")
)
