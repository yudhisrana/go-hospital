package usecase

import (
	"context"
	"time"

	"github.com/yudhisrana/go-hospital/internal/application/patient/dto"
	"github.com/yudhisrana/go-hospital/internal/domain/patient/entity"
	"github.com/yudhisrana/go-hospital/internal/domain/patient/repository"
)

type PatientCreateUsecase struct {
	repository.CreatePatientRepository
}

type CreatePatient struct {
	patientRepo PatientCreateUsecase
}

func NewCreatePatientUseCase(patientRepo PatientCreateUsecase) *CreatePatient {
	return &CreatePatient{patientRepo: patientRepo}
}

func (u *CreatePatient) Execute(ctx context.Context, req *dto.RequestPayload) (*dto.ResponsePayload, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	newPatient, err := entity.NewPatient(req.NIK, req.Name, req.Age, req.Gender, req.BirthDate, req.Address, req.PhoneNumber)
	if err != nil {
		return nil, err
	}

	if err := u.patientRepo.Save(ctx, newPatient); err != nil {
		return nil, err
	}

	response := dto.NewResponsePayloadPost(newPatient)

	return response, nil
}
