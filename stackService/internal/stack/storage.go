package stack

import (
	"context"

	"stackService/dto"
)

type Storage interface {
	Create(ctx context.Context, data dto.StackData) (dto.StackData, error)
	GetAndDelete(ctx context.Context) (dto.StackData, error)
}
