package senddataset

import (
	"context"
	"log"
	"time"

	"github.com/johann-vu/iot-scenario/internal/domain"
)

type Service interface {
	Execute(ctx context.Context) error
}

type MeasurementGenerator func() float64

type service struct {
	serviceId string
	generate  MeasurementGenerator
	sender    domain.DatasetSender
}

// Execute implements Service.
func (s *service) Execute(ctx context.Context) error {

	d := domain.Dataset{
		SensorID:  s.serviceId,
		Timestamp: time.Now(),
		Value:     s.generate(),
	}

	err := s.sender.Send(ctx, d)
	if err != nil {
		log.Printf("error while sending data: %v\n", err)
		return err
	}

	log.Println("data was sent successfully")
	return nil
}

func NewService(serviceId string, generateFunc MeasurementGenerator, sender domain.DatasetSender) Service {

	return &service{
		serviceId: serviceId,
		generate:  generateFunc,
		sender:    sender,
	}
}
