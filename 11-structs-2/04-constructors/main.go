package main

import (
	"errors"
	"fmt"
	"time"
)

type DateRange struct {
	Start time.Time
	End   time.Time
}

func (d DateRange) Hours() float64 {
	return d.End.Sub(d.Start).Hours()
}

func main() {
	//lifetime := DateRange{
	//	Start: time.Date(1814, 12, 10, 0, 0, 0, 0, time.UTC),
	//	End:   time.Date(1851, 11, 27, 0, 0, 0, 0, time.UTC),
	//}

	lifetime, err := NewDateRange(time.Date(1814, 12, 10, 0, 0, 0, 0, time.UTC), time.Date(1851, 11, 27, 0, 0, 0, 0, time.UTC))

	if err == nil {
		fmt.Println(lifetime.Hours())
	}

	//travelInTime := DateRange{
	//	Start: time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
	//	End:   time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
	//}

	travelInTime, err := NewDateRange(
		time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
		time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
	)

	if err == nil {
		fmt.Println(travelInTime.Hours())
	}
}

func NewDateRange(start, end time.Time) (DateRange, error) {
	if time.Time.IsZero(start) || time.Time.IsZero(end) {
		return DateRange{}, errors.New("invalid time")
	}

	if time.Time.Before(end, start) {
		return DateRange{}, errors.New("can't time travel")
	}

	return DateRange{
		start,
		end,
	}, nil
}
