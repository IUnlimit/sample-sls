package model

import "time"

type Config struct {
	Log      *Log          `yaml:"log"`
	Interval time.Duration `yaml:"interval"`
	App      []*App        `yaml:"app"`
}

type Log struct {
	ForceNew bool          `yaml:"force-new,omitempty"`
	Level    string        `yaml:"level,omitempty"`
	Aging    time.Duration `yaml:"aging,omitempty"`
	Colorful bool          `yaml:"colorful,omitempty"`
}

type App struct {
	Converter *Converter `yaml:"converter"`
	Formater  *Formater  `yaml:"formater"`
	Publisher *Publisher `yaml:"publisher,omitempty"` // pipeline=true时, 无需配置 publisher
	Pipeline  bool       `yaml:"pipeline,omitempty"`
}

// Converter 转换器配置
type Converter struct {
	Type     string `yaml:"type"`
	Location string `yaml:"location,omitempty"` // 仅 pipeline=false 使用
	Regex    *Regex `yaml:"regex,omitempty"`
}

// Regex Converter 转换器子配置
type Regex struct {
	Pattern string `yaml:"pattern,omitempty"`
}

// Formater 格式化配置
type Formater struct {
	Type    string   `yaml:"type"`
	Journal *Journal `yaml:"journal,omitempty"`
}

type Journal struct {
	Prefix string `yaml:"prefix,omitempty"`
}

type Publisher struct {
	Type string `yaml:"type"`
}
