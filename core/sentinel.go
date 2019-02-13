package core

import (
	"fmt"
	"github.com/pedrolopesme/sentinel/client"
	"github.com/satori/go.uuid"
)

type Sentinel interface {
	// GetId returns a unique identifier to the sentinel
	GetId() string

	// Run puts sentinel to run and returns its execution Id and an error
	Run(stockProvider client.StockProvider) (string, error)

	// Kill stops a Sentinel
	Kill() error
}

// TODO extract alphaVantageKey from here.
// TODO decouple sentinel logic from http consume logic.
// TODO Alphavantage deserves its own client.
type StockSentinel struct {
	id       string
	config   *SentinelConfig
	schedule *Schedule
}

// GetId returns a unique identifier to the sentinel
// TODO add tests
func (s *StockSentinel) GetId() string {
	return s.id
}

// GetId returns a unique identifier to the sentinel
// TODO add tests
// TODO add log
// TODO extract all AlphaVantageKey retrieval to its own client
func (s *StockSentinel) Run(stockProvider client.StockProvider) (string, error) {
	var executionId = uuid.Must(uuid.NewV4()).String()
	fmt.Println("Running StockSentinel ", s.GetId(), " - execution ", executionId)

	_, err := stockProvider.GetStocks(s.schedule.Stock, s.schedule.TimeFrame)
	if err != nil {
		fmt.Println("Cant get stocks due to", err.Error())
	}

	return executionId, nil
}

// Kill stops a Sentinel
// TODO add tests
func (s *StockSentinel) Kill() error {
	return nil
}

// NewSentinel
// TODO add tests
func NewStockSentinel(config *SentinelConfig, schedule *Schedule) (sentinel *StockSentinel) {
	return &StockSentinel{
		id:       uuid.Must(uuid.NewV4()).String(),
		schedule: schedule,
		config:   config,
	}
}
