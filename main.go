package main

import (
	"github.com/IUnlimit/sample-sls/internal/dataflow"
	"time"
)

func main() {
	p := dataflow.NewPipeline()
	fixedWindow := dataflow.NewFixedWindow(p, 2*time.Second)
	//dataflow.SequenceOf(
	//	dataflow.RepeatUnit(
	//		dataflow.AtPeriod(),
	//		dataflow.AtWatermark(),
	//	),
	//	dataflow.RepeatOn(dataflow.AtWatermark()),
	//)
	println(fixedWindow)
}
