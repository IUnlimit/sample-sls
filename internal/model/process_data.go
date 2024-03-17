package model

type ProcessData struct {
	// Watermarks labeled `Processing Time` which is a lower bound
	Watermarks int64
}
