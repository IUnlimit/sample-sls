package dataflow

import "github.com/golang-collections/collections/set"

// ParDo parse input
type ParDo interface {
}

// GroupByKey parse output
type GroupByKey interface {
	// GroupByKey groups (value, window) tuples by key.
	GroupByKey()
}

type GroupByKeyAndWindow struct {
}

// AssignWindows assigns the element to zero or more windows
func (gbk *GroupByKeyAndWindow) AssignWindows(datum string) {
}

// DropTimestamps drops element timestamps, as only the window is relevant from here on out
func (gbk *GroupByKeyAndWindow) DropTimestamps() {
}

func (gbk *GroupByKeyAndWindow) GroupByKey() {
}

// MergeWindows merges windows at grouping time
func (gbk *GroupByKeyAndWindow) MergeWindows(windows set.Set) {
}

// GroupAlsoByWindow for each key, groups values by window
func (gbk *GroupByKeyAndWindow) GroupAlsoByWindow(windows set.Set) {
}

// ExpandToElements expands per-key, per-window groups of values into (key, value, eventtime, window)tuples, with new per-window timestamps
func (gbk *GroupByKeyAndWindow) ExpandToElements(windows set.Set) {
}
