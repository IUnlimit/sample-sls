package entry

import (
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/graph/mtime"
)

func Init() {
	parser = NewDateTimeRegexParser()
}

// ParseLogEntryDoFn parse log str to LogEntry
func ParseLogEntryDoFn(element string, emit func(*LogEntry)) {
	entry := &LogEntry{Topic: &Topic{Content: element}}
	emit(entry)
}

// AddTimestampDoFn extracts an event time from a LogEntry.
func AddTimestampDoFn(element *LogEntry, emit func(beam.EventTime, *LogEntry)) error {
	et, err := parser.extractEventTime(element)
	if err != nil {
		return err
	}
	emit(mtime.FromTime(*et), element)
	return nil
}
