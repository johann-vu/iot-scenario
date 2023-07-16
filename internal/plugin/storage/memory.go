package storage

import (
	"context"
	"time"

	"github.com/johann-vu/iot-scenario/internal/domain"
)

func NewMemoryDatasetRepository() domain.DatasetRepository {
	return &memory{data: []domain.Dataset{}}
}

type memory struct {
	data []domain.Dataset
}

// Add implements domain.DatasetRepository.
func (dr *memory) Add(ctx context.Context, d domain.Dataset) error {
	dr.data = append(dr.data, d)
	return nil
}

// Get implements domain.DatasetRepository.
func (dr *memory) Get(ctx context.Context, from time.Time, to time.Time) ([]domain.Dataset, error) {
	result := []domain.Dataset{}

	for _, d := range dr.data {
		if isInTimerange(d, from, to) {
			result = append(result, d)
		}
	}

	return result, nil
}

func isInTimerange(d domain.Dataset, from time.Time, to time.Time) bool {
	return d.Timestamp.After(from) && d.Timestamp.Before(to)
}
