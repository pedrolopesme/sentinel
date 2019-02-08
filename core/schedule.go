package core

// TODO Stock should be a type (?)
// TODO TimeFrame should be a type (?)
type Schedule struct {
	Stock     string
	TimeFrame string
}

// TODO add validations
// TODO add tests
func NewSchedule(stock string, timeFrame string) *Schedule {
	return &Schedule{
		Stock:     stock,
		TimeFrame: timeFrame,
	}
}
