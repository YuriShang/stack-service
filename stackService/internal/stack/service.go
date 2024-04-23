package stack

import (
	"context"

	d "stackService/dto"
	"stackService/pkg/logging"
)

var _ Service = &service{}

type service struct {
	storage Storage
	logger  logging.Logger
}

func NewService(stackStorage Storage, logger logging.Logger) (Service, error) {
	return &service{
		storage: stackStorage,
		logger:  logger,
	}, nil
}

type Service interface {
	Push(ctx context.Context, dto d.StackData) (d.StackData, error)
	Pop(ctx context.Context) (d.StackData, error)
}

func (s service) Push(ctx context.Context, dto d.StackData) (d.StackData, error) {
	createData := d.NewData(dto)
	data, err := s.storage.Create(ctx, createData)
	return data, err
}

func (s service) Pop(ctx context.Context) (d.StackData, error) {
	data, err := s.storage.GetAndDelete(ctx)
	return data, err
}
