package senddataset

import (
	"math/rand"
	"time"
)

func NewRandomGenerator(max, min float64) MeasurementGenerator {
	return func() float64 {
		rand.Seed(time.Now().UnixNano())
		return (max-min)*rand.Float64() + min
	}
}
