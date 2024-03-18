package dataflow

import "time"

// Window
// Windowing determines where in event time data are grouped together for processing
type Window interface {
}

type FixedWindow struct {
	p *Pipeline
	d time.Duration
}

func NewFixedWindow(pipeline *Pipeline, duration time.Duration) *FixedWindow {
	return &FixedWindow{
		p: pipeline,
		d: duration,
	}
}
