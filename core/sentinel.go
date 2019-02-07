package core

import (
	"fmt"
	"github.com/satori/go.uuid"
)

type Sentinel interface {
	// GetId returns a unique identifier to the sentinel
	GetId() string

	// Run puts sentinel to run and returns its execution Id and an error
	Run() (string, error)

	// Kill stops a Sentinel
	Kill() error
}

type StockSentinel struct {
	id string
}

// GetId returns a unique identifier to the sentinel
func (s *StockSentinel) GetId() string {
	return s.id
}

// GetId returns a unique identifier to the sentinel
func (s *StockSentinel) Run() (string, error) {
	var executionId = uuid.Must(uuid.NewV4()).String()
	fmt.Println("Running StockSentinel ", executionId)
	return executionId, nil
}

// Kill stops a Sentinel
func (s *StockSentinel) Kill() error {
	return nil
}

// NewSentinel
func NewStockSentinel() (sentinel *StockSentinel) {
	return &StockSentinel{
		id: uuid.Must(uuid.NewV4()).String(),
	}
}
