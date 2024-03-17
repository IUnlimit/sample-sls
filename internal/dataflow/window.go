package dataflow

import "time"

// Window
// Windowing determines where in event time data are grouped together for processing
type Window interface {
	DoSth()
}

type Sessions struct {
	// todo
	// current time
	duration time.Duration
}

func WithGapDuration(d time.Duration) *Sessions {
	return &Sessions{
		duration: d,
	}
}

func (s *Sessions) DoSth() {

}
