package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/yudhisrana/go-hospital/internal/domain/visit/entity"
)

type CreateVisitRepository interface {
	Save(ctx context.Context, visit *entity.Visit) error
}

type FindVisitRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Visit, error)
	FindAll(ctx context.Context) ([]*entity.Visit, error)
	FindByStatus(ctx context.Context, status string) ([]*entity.Visit, error)
}

type UpdateVisitRepository interface {
	Update(ctx context.Context, visit *entity.Visit) error
}

type GetNextQueueNumberRepository interface {
	GetNextQueueNumber(ctx context.Context) (int, error)
}
