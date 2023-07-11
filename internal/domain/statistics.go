package domain

type Statistics struct {
	Average float32
}

type StatisticsCalculator interface {
	CalculateStatistics(d []Dataset) (Statistics, error)
}
