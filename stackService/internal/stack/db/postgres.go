package db

import (
	"context"

	"stackService/dto"
	"stackService/internal/stack"
	"stackService/pkg/logging"
	"stackService/pkg/postgres"
)

var _ stack.Storage = &db{}

type db struct {
	client *postgres.Client
	logger logging.Logger
}

func NewStorage(client *postgres.Client, logger logging.Logger) stack.Storage {
	return &db{
		client: client,
		logger: logger,
	}
}

func (s *db) GetAndDelete(ctx context.Context) (dto.StackData, error) {
	data, err := s.client.GetAndDelete(ctx)
	return data, err
}

func (s *db) Create(ctx context.Context, data dto.StackData) (dto.StackData, error) {
	data, err := s.client.Create(ctx, data)
	return data, err
}
