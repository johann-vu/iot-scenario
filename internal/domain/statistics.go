package domain

type Statistics struct {
	Average float64
}

type StatisticsCalculator interface {
	CalculateStatistics(d []Dataset) (Statistics, error)
}
