package main

import (
	"github.com/IUnlimit/sample-sls/cmd/sls"
	"github.com/IUnlimit/sample-sls/internal/conf"
	"github.com/IUnlimit/sample-sls/internal/logger"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/local"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/runners/direct"
)

func main() {
	conf.Init()
	logger.Init()
	sls.ServeRESTFul()
	//beam.Init()
	//entry.Init()
	//p, s := beam.NewPipelineWithRoot()
	//
	//lines := textio.Read(s, "README.md")
	//moreLines := textio.Read(s, "LICENSE")
	//all := beam.Flatten(s, lines, moreLines)
	//
	//logEntries := beam.ParDo(s, entry.ParseLogEntryDoFn, all)
	//timestampedLines := beam.ParDo(s, entry.AddTimestampDoFn, logEntries)
	//
	//triggers := make([]trigger.Trigger, 0)
	//triggers = append(triggers, trigger.AfterProcessingTime().PlusDelay(60*time.Second))
	//triggers = append(triggers, trigger.Repeat(trigger.AfterCount(10)))
	//afterAny := trigger.AfterAny(triggers)
	//
	//windowedLines := beam.WindowInto(
	//	s,
	//	window.NewGlobalWindows(),
	//	timestampedLines,
	//	beam.Trigger(afterAny),
	//)
	//
	//wordRegexp := regexp.MustCompile(`[a-zA-Z]+('[a-z])?`)
	//words := beam.ParDo(s, func(line string, emit func(t string)) {
	//	for _, word := range wordRegexp.FindAllString(line, -1) {
	//		emit(word)
	//	}
	//}, windowedLines)
	//formatted := beam.ParDo(s, strings.ToUpper, words)
	//
	//textio.Write(s, "output/1.txt", formatted)
	//
	//if _, err := beam.Run(context.Background(), "direct", p); err != nil {
	//	fmt.Printf("Pipeline failed: %v", err)
	//}
}

//// Use the DoFn with ParDo as normal.
//stampedLogs := beam.ParDo(s, AddTimestampDoFn, unstampedLogs)
