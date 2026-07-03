package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/yudhisrana/go-hospital/internal/domain/patient/entity"
)

type FindAllPatientRepository interface {
	FindAll(ctx context.Context) ([]*entity.Patient, error)
}

type FindPatientByIDRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Patient, error)
}

type CreatePatientRepository interface {
	Save(ctx context.Context, patient *entity.Patient) error
}

type UpdatePatientRepository interface {
	Update(ctx context.Context, patient *entity.Patient) error
}

type DeletePatientRepository interface {
	Delete(ctx context.Context, id uuid.UUID) error
}
