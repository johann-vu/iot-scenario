package storedataset

import (
	"context"
	"log"

	"github.com/johann-vu/iot-scenario/internal/domain"
)

type Service interface {
	Execute(ctx context.Context, d domain.Dataset) error
}

type service struct {
	repository domain.DatasetRepository
}

// Execute implements Service.
func (srv *service) Execute(ctx context.Context, d domain.Dataset) error {
	err := srv.repository.Add(ctx, d)

	if err != nil {
		log.Printf("error while storing dataset: %v\n", err)
		return err
	}

	log.Printf("data from sensor %q was stored successfully", d.SensorID)
	return nil
}

func NewService(repository domain.DatasetRepository) Service {
	return &service{repository: repository}
}
