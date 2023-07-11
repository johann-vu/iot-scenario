package storedataset

import (
	"context"

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
	return srv.repository.Add(ctx, d)
}

func NewService(repository domain.DatasetRepository) Service {
	return &service{repository: repository}
}
