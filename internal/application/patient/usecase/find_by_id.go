package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/yudhisrana/go-hospital/internal/application/patient/dto"
	"github.com/yudhisrana/go-hospital/internal/domain/patient/repository"
)

type PatientFindByIDUsecase struct {
	repository.FindPatientByIDRepository
}

type FindPatientByID struct {
	patientRepo PatientFindByIDUsecase
}

func NewFindPatientByIDUseCase(patientRepo PatientFindByIDUsecase) *FindPatientByID {
	return &FindPatientByID{patientRepo: patientRepo}
}

func (u *FindPatientByID) Execute(ctx context.Context, id uuid.UUID) (*dto.ResponsePayload, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	patient, err := u.patientRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if patient == nil {
		return nil, nil
	}

	return dto.NewResponsePayloadPost(patient), nil
}



