package http

import (
	"time"

	"github.com/johann-vu/iot-scenario/internal/domain"
)

type StatisticsDTO struct {
	Average           float64      `json:"average"`
	StandardDeviation float64      `json:"standardDeviation"`
	Minimum           DatasetDTO   `json:"minimum"`
	Maximum           DatasetDTO   `json:"maximum"`
	Count             int          `json:"count"`
	Slope             float64      `json:"estimatedSlope"`
	Recent            []DatasetDTO `json:"recentValues"`
}

type DatasetDTO struct {
	SensorID      string  `json:"sensorId" validate:"required"`
	UnixTimestamp int64   `json:"unixTimestamp" validate:"required,gte=0"`
	Value         float64 `json:"value" validate:"required"`
}

func (d DatasetDTO) ToDomain() domain.Dataset {
	return domain.Dataset{
		SensorID:  d.SensorID,
		Timestamp: time.Unix(d.UnixTimestamp, 0),
		Value:     d.Value,
	}
}
