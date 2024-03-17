package model

// MetaData log meta data
type MetaData struct {
	// log Source 127.0.0.1
	Source string `json:"__source__"`
	Tag    *Tag   `json:"__tag__"`
	Topic  *Topic `json:"__topic__"`
}

// Tag log tag
type Tag struct {
	// Hostname iunlimit
	Hostname string `json:"__hostname__,omitempty"`
	// log storage Path /var/log/exec/file.log
	Path string `json:"__path__,omitempty"`
}

// Topic log topic
type Topic struct {
	// log Content (with `Event Time`)
	Content string `json:"content"`
}
