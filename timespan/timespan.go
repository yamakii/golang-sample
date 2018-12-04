package timespan

import (
	"errors"
	"time"
)

type TimeSpan struct {
	start time.Time
	end   time.Time
}

func New(start, end time.Time) (TimeSpan, error) {
	if end.Before(start) {
		return TimeSpan{}, errors.New("endはstartよりも前ではいけません")
	}
	return TimeSpan{start: start, end: end}, nil
}
