package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/yudhisrana/go-hospital/internal/domain/patient/repository"
)

type DeletePatientUseCase struct {
	repository.FindPatientByIDRepository
	repository.DeletePatientRepository
}

type DeletePatient struct {
	patientRepo DeletePatientUseCase
}

func NewDeletePatientUseCase(patientRepo DeletePatientUseCase) *DeletePatient {
	return &DeletePatient{patientRepo: patientRepo}
}

func (u *DeletePatient) Execute(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Check if patient exists
	_, err := u.patientRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return u.patientRepo.Delete(ctx, id)
}


