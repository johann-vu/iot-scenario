package calculatestatistics

import (
	"math"

	"github.com/johann-vu/iot-scenario/internal/domain"
	"gonum.org/v1/gonum/stat"
)

func CalculateExtremeValues(data *[]domain.Dataset, result *domain.Statistics) {
	max := (*data)[0]
	min := (*data)[0]
	for _, d := range *data {
		if d.Value > max.Value {
			max = d
			continue
		}
		if d.Value < min.Value {
			min = d
		}
	}

	result.Maximum = max
	result.Minimum = min
}

func CalculateStandardDeviation(data *[]domain.Dataset, result *domain.Statistics) {

	var sum float64
	for _, d := range *data {
		sum += d.Value
	}

	n := float64(len(*data))

	mean := sum / n

	sumSquaredDiff := 0.0
	for _, d := range *data {
		sumSquaredDiff += math.Pow(d.Value-mean, 2)
	}

	stdDev := math.Sqrt(sumSquaredDiff / n)

	result.StandardDeviation = stdDev
}

func CalculateCount(data *[]domain.Dataset, result *domain.Statistics) {
	result.Count = len(*data)
}

func CalculateAverage(data *[]domain.Dataset, result *domain.Statistics) {
	n := float64(len(*data))
	if n < 2 {
		return
	}

	var sum float64

	for _, d := range *data {
		sum += d.Value
	}

	result.Average = sum / n
}

func CalculateTrend(data *[]domain.Dataset, result *domain.Statistics) {
	n := float64(len(*data))
	if n < 2 {
		return
	}

	var sumX, sumY, sumXY, sumX2 float64

	for _, d := range *data {
		t := float64(d.Timestamp.Unix())
		sumX += t
		sumY += d.Value
		sumXY += t * d.Value
		sumX2 += math.Pow(t, 2)
	}

	slope := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)

	result.Trend = slope
}

func CalculateLinearRegression(data *[]domain.Dataset, result *domain.Statistics) {
	var (
		xs      = make([]float64, len(*data))
		ys      = make([]float64, len(*data))
		weights []float64
	)

	for i := range xs {
		xs[i] = float64((*data)[i].Timestamp.Unix())
		ys[i] = (*data)[i].Value
	}

	offset, slope := stat.LinearRegression(xs, ys, weights, false)
	result.LinearRegression = domain.LinearRegression{
		Offset: offset, Slope: slope,
	}
}
