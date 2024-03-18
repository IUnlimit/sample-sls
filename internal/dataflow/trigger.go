package dataflow

// Trigger allow the specification of when to trigger the output results for a given window
// It's the complementary to the windowing model
// Triggering determines when in processing time the results of groupings are emitted as panes
type Trigger struct {
}

// RefinementMode controls how multiple panes for the same window relate to each other
type RefinementMode int8

const (
	// Discarding Upon triggering, window contents are discarded,
	// and later results bear no relation to previous results
	Discarding RefinementMode = 1

	// Accumulating Upon triggering, window contents are left intact in persistent state,
	// and later results become a refinement of previous results
	Accumulating RefinementMode = 2

	// AccumulatingAndRetracting Upon triggering, inaddition to the Accumulating semantics,
	// a copy of the emitted value is also stored in persistent state
	AccumulatingAndRetracting RefinementMode = 3
)

// Sequence is a trigger combiner used to combine multiple triggers in a specific order
// With Sequence, you can define the execution order of triggers so that the output of
// one trigger becomes the input of the next trigger.
type Sequence struct {
	repeats []*Repeat
}

func SequenceOf(repeats ...*Repeat) *Sequence {
	return &Sequence{
		repeats: repeats,
	}
}

// Repeat is a trigger combiner that is used to repeatedly execute specific triggers according to certain rules
// Through Repeat, you can specify the number of times or conditions for a trigger to be executed repeatedly,
// thereby controlling the execution logic of data processing operations.
type Repeat struct {
	// There is an "or" relationship between conditions
	conds []*Condition
}

func RepeatOn(condition *Condition) *Repeat {
	conditions := make([]*Condition, 0)
	conditions = append(conditions, condition)
	return &Repeat{
		conds: conditions,
	}
}

func RepeatUnit(conditions ...*Condition) *Repeat {
	return &Repeat{
		conds: conditions,
	}
}
