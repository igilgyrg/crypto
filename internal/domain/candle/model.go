package candle

import "time"

type Candle struct {
	Max         float64   `json:"max" bson:"max"`
	Min         float64   `json:"min" bson:"min"`
	Open        float64   `json:"open" bson:"open"`
	Closed      float64   `json:"closed" bson:"closed"`
	UpperShadow float64   `json:"upper_shadow,omitempty" bson:"upper_shadow"`
	LowerShadow float64   `json:"lower_shadow,omitempty" bson:"lower_shadow"`
	Datetime    time.Time `json:"datetime" bson:"datetime"`
}
