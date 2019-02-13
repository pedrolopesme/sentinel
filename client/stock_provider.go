package client

import (
	"github.com/pedrolopesme/sentinel/models"
	"time"
)

// StockProvider represents the minimum interface that any
// stock provider should have.
type StockProvider interface {
	GetStocks(stock string, timeFrame string) (map[time.Time]models.StockTier, error)
}
