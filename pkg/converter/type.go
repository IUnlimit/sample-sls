package converter

const (
	Json = "Json"
	// Logfmt https://pkg.go.dev/github.com/kr/logfmt
	Logfmt = "Logfmt"
	// Regex with PCRE standard
	Regex = "Regex"
)

type Converter interface {
	extract() error
}
