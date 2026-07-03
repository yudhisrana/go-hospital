package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/yudhisrana/go-hospital/internal/application/patient/dto"
	"github.com/yudhisrana/go-hospital/internal/domain/patient/repository"
)

type PatientUpdateUsecase struct {
	repository.UpdatePatientRepository
	repository.FindPatientByIDRepository
}

type UpdatePatient struct {
	patientRepo PatientUpdateUsecase
}

func NewUpdatePatientUseCase(patientRepo PatientUpdateUsecase) *UpdatePatient {
	return &UpdatePatient{patientRepo: patientRepo}
}

func (u *UpdatePatient) Execute(ctx context.Context, id uuid.UUID, req *dto.RequestPayload) (*dto.ResponsePayload, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	p, err := u.patientRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, nil
	}

	if err := p.UpdatePatient(req.NIK, req.Name, req.Age, req.Gender, req.BirthDate, req.Address, req.PhoneNumber); err != nil {
		return nil, err
	}

	if err := u.patientRepo.Update(ctx, p); err != nil {
		return nil, err
	}

	response := dto.NewResponsePayloadPost(p)
	return response, nil
}

