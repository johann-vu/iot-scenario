package domain

import (
	"context"
	"time"
)

type Dataset struct {
	SensorID  string
	Timestamp time.Time
	Value     float64
}

type DatasetRepository interface {
	Add(ctx context.Context, d Dataset) error
	Get(ctx context.Context, from time.Time, to time.Time) ([]Dataset, error)
}

type DatasetSender interface {
	Send(ctx context.Context, d Dataset) error
}
