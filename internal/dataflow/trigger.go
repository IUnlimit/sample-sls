package dataflow

// Trigger allow the specification of when to trigger the output results for a given window
// It's the complementary to the windowing model
// Triggering determines when in processing time the results of groupings are emitted as panes
type Trigger interface {
}

type RefinementMode int8
