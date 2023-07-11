package generatedataset

import (
	"math/rand"
	"time"

	"github.com/johann-vu/iot-scenario/internal/domain"
)

type Service interface {
	Execute() (domain.Dataset, error)
}

type service struct {
	serviceId string
	maxValue  float32
	minValue  float32
}

// Execute implements Service.
func (s *service) Execute() (domain.Dataset, error) {
	return domain.Dataset{
		SensorID:  s.serviceId,
		Timestamp: time.Now(),
		Value:     s.CalculateNextValue(),
	}, nil
}

func (s *service) CalculateNextValue() float32 {
	rand.Seed(time.Now().UnixNano())
	return (s.maxValue-s.minValue)*rand.Float32() + s.minValue
}

func NewService(serviceId string, min, max float32) Service {
	if min > max {
		panic("invalid values for min and max")
	}

	return &service{
		serviceId: serviceId,
		maxValue:  max,
		minValue:  min,
	}
}
