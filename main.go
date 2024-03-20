package main

import (
	"context"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/graph/mtime"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/graph/window"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/graph/window/trigger"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/runners/direct"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/local"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/textio"
)

type addTimestampFn struct {
	Min beam.EventTime `json:"min"`
}

func (f *addTimestampFn) ProcessElement(x beam.X) (beam.EventTime, beam.X) {
	timestamp := f.Min.Add(time.Duration(rand.Int63n(1 * time.Hour.Nanoseconds())))
	return timestamp, x
}

func main() {
	beam.Init()
	p, s := beam.NewPipelineWithRoot()

	lines := textio.Read(s, "D:\\code\\go\\sample-sls\\README.md")
	moreLines := textio.Read(s, "D:\\code\\go\\sample-sls\\LICENSE")
	all := beam.Flatten(s, lines, moreLines)
	timestampedLines := beam.ParDo(s, &addTimestampFn{Min: mtime.Now()}, all)

	triggers := make([]trigger.Trigger, 0)
	triggers = append(triggers, trigger.AfterProcessingTime().PlusDelay(60*time.Second))
	triggers = append(triggers, trigger.Repeat(trigger.AfterCount(10)))
	afterAny := trigger.AfterAny(triggers)

	windowedLines := beam.WindowInto(
		s,
		window.NewGlobalWindows(),
		timestampedLines,
		beam.Trigger(afterAny),
	)

	wordRegexp := regexp.MustCompile(`[a-zA-Z]+('[a-z])?`)
	words := beam.ParDo(s, func(line string, emit func(t string)) {
		for _, word := range wordRegexp.FindAllString(line, -1) {
			emit(word)
		}
	}, windowedLines)
	formatted := beam.ParDo(s, strings.ToUpper, words)

	textio.Write(s, "D:\\code\\go\\sample-sls\\output", formatted)

	if _, err := beam.Run(context.Background(), "direct", p); err != nil {
		fmt.Printf("Pipeline failed: %v", err)
	}
}

// 从文件读取日志需要手动转换timestamp
// // AddTimestampDoFn extracts an event time from a LogEntry.
//func AddTimestampDoFn(element LogEntry, emit func(beam.EventTime, LogEntry)) {
//	et := extractEventTime(element)
//	// Defining an emitter with beam.EventTime as the first parameter
//	// allows the DoFn to set the event time for the emitted element.
//	emit(mtime.FromTime(et), element)
//}
//
//// Use the DoFn with ParDo as normal.
//stampedLogs := beam.ParDo(s, AddTimestampDoFn, unstampedLogs)

//     // 将数据分配到固定时间窗口，并设置每个窗口最大元素数量为 3
//    windowedData := beam.WindowInto(s, beam.NewGlobalWindows(), data)
//    windowedData = beam.ParDo(s, func(i int) (beam.Window, int) {
//        return beam.GlobalWindow{}, i
//    }, windowedData)
//
//    // 使用 Count 触发器来限制每个窗口的元素数量
//    triggeredData := beam.Trigger(windowedData, periodic.Trigger{
//        Trigger: periodic.CountTrigger{MaxElement: 3},
//    })
