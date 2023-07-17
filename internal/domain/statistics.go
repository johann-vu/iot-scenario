package domain

type Statistics struct {
	Average           float64
	StandardDeviation float64
	Minimum           Dataset
	Maximum           Dataset
	Trend             float64
	Count             int
	LinearRegression  LinearRegression
}

type LinearRegression struct {
	Offset float64
	Slope  float64
}
