package candle

import "time"

type CreateCandle struct {
	Max      float64   `json:"max"`
	Min      float64   `json:"min"`
	Open     float64   `json:"open"`
	Closed   float64   `json:"closed"`
	Datetime time.Time `json:"datetime"`
}
