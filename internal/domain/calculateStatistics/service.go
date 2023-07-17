package calculatestatistics

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/johann-vu/iot-scenario/internal/domain"
)

type service struct {
	repo domain.DatasetRepository
}

type calculateFunc func(*[]domain.Dataset, *domain.Statistics)

var calculations []calculateFunc = []calculateFunc{
	CalculateAverage, CalculateCount, CalculateExtremeValues, CalculateStandardDeviation, CalculateLinearRegression,
}

// Execute implements Service.
func (srv *service) Execute(ctx context.Context, from, to time.Time) (domain.Statistics, error) {

	defer func(start time.Time) {
		log.Printf("calculated statistics in %dms", time.Since(start).Milliseconds())
	}(time.Now())

	data, err := srv.repo.Get(ctx, from, to)
	if err != nil {
		log.Printf("error while loading data: %s", err)
		return domain.Statistics{}, err
	}

	if len(data) < 2 {
		return domain.Statistics{}, fmt.Errorf("not enough data (%d)", len(data))
	}

	result := domain.Statistics{}

	for _, cf := range calculations {
		cf(&data, &result)
	}

	return result, nil
}

type Service interface {
	Execute(ctx context.Context, from, to time.Time) (domain.Statistics, error)
}

func NewService(repo domain.DatasetRepository) Service {
	return &service{repo: repo}
}
