package dataflow

type Condition struct {
}

func AtWatermark() *Condition {
	return nil
}

func AtPeriod() *Condition {
	return nil
}

func AtPercentileWatermark() *Condition {
	return nil
}
