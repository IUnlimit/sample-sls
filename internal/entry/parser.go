package entry

import (
	"regexp"
	"time"
)

var parser *RegexParser

type Parser interface {
	extractEventTime(element LogEntry) time.Time
}

type RegexParser struct {
	// time string regex
	regex *regexp.Regexp
	// time parse layout
	layout string
}

func NewDateTimeRegexParser() *RegexParser {
	return &RegexParser{
		regex:  DateTimeRegex,
		layout: DateTimeLayout,
	}
}

func NewCustomRegexParser(regx string) *RegexParser {
	return &RegexParser{
		regex: regexp.MustCompile(regx),
	}
}

func (r *RegexParser) extractEventTime(entry *LogEntry) (*time.Time, error) {
	content := entry.Topic.Content
	if dateTime := r.regex.FindString(content); len(dateTime) != 0 {
		timeObj, err := time.Parse(r.layout, dateTime)
		if err != nil {
			return nil, err
		}
		return &timeObj, nil
	}
	return nil, nil
}
