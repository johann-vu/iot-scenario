package senddataset

import (
	"math"
	"math/rand"
	"time"
)

func NewRandomGenerator(max, min float64) MeasurementGenerator {
	return func() float64 {
		rand.Seed(time.Now().UnixNano())
		return (max-min)*rand.Float64() + min
	}
}

func NewWaveGenerator(max, min, frequency float64) MeasurementGenerator {
	i := 0
	return func() float64 {
		i++
		return (max-min)*((math.Sin(float64(i)*(frequency))/2)+0.5) + min
	}
}

func NewLinearGenerator(min, slope float64) MeasurementGenerator {
	i := 0
	return func() float64 {
		i++
		return slope*float64(i) + min
	}
}
