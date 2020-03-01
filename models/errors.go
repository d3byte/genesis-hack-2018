package models

import "errors"

var (
	ErrNotFound       = errors.New("request is not found")
	ErrTimeIncorrect  = errors.New("error: end time field before start time")
	ErrTimeEquals     = errors.New("end and start times is equals")
	ErrEventsNotFound = errors.New("no upcoming events found")
	ErrTimeDuration   = errors.New("duration must be 2 hours")
)
