package core

type Schedule struct {
	Stock     string
	TimeFrame string
}

func NewSchedule(stock string, timeFrame string) *Schedule {
	return &Schedule{
		Stock:     stock,
		TimeFrame: timeFrame,
	}
}
