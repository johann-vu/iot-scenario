package domain

type Statistics struct {
	Average           float64
	StandardDeviation float64
	Minimum           Dataset
	Maximum           Dataset
	Count             int
	LinearRegression  LinearRegression
	Recent            []Dataset
}

type LinearRegression struct {
	Offset float64
	Slope  float64
}
