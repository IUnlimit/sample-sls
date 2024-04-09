package model

import "time"

type Config struct {
	Log *Log           `yaml:"log"`
	API *API           `yaml:"api"`
	ES  *Elasticsearch `yaml:"es"`
}

type Log struct {
	ForceNew bool          `yaml:"force-new,omitempty"`
	Level    string        `yaml:"level,omitempty"`
	Aging    time.Duration `yaml:"aging,omitempty"`
	Colorful bool          `yaml:"colorful,omitempty"`
}

type API struct {
	Host string `yaml:"host,omitempty"`
	Port int64  `yaml:"port,omitempty"`
}

type Elasticsearch struct {
	Url string `yaml:"url,omitempty"`
}
