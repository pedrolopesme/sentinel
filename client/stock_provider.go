package client

import (
	"github.com/pedrolopesme/sentinel/models"
	"time"
)

// StocksByTime represents a group of stocks
type StocksByTime map[time.Time]models.StockTier

// StockProvider represents the minimum interface that any
// stock provider should have.
type StockProvider interface {
	GetName() string
	GetStocks(stock string, timeFrame string) (*StocksByTime, error)
}
