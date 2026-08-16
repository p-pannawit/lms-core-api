package health

import (
	"context"
	"time"
)

const checkTimeout = 2 * time.Second

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) Check(ctx context.Context) error {
	checkContext, cancel := context.WithTimeout(ctx, checkTimeout)
	defer cancel()

	return service.repository.Ping(checkContext)
}
