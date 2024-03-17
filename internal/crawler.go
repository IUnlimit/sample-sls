package internal

// Crawler is responsible for obtaining a single input
type Crawler interface {
	take() string
}
