package core

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestNewScheduleShouldReturnAProperSchedule(t *testing.T) {
	var (
		expectedStock     = "MyStock"
		expectedTimeFrame = "TimeFrame"
		assert            = assert2.New(t)
	)

	newSchedule := NewSchedule(expectedStock, expectedTimeFrame)
	assert.NotNil(newSchedule)
	assert.Equal(expectedStock, newSchedule.Stock)
	assert.Equal(expectedTimeFrame, newSchedule.TimeFrame)
}
