package usecase

import (
	"context"
	"time"

	"github.com/yudhisrana/go-hospital/internal/application/patient/dto"
	"github.com/yudhisrana/go-hospital/internal/domain/patient/repository"
)

type PatientFindAllUsecase struct {
	repository.FindAllPatientRepository
}

type FindAllPatient struct {
	patientRepo PatientFindAllUsecase
}

func NewFindAllPatientUseCase(patientRepo PatientFindAllUsecase) *FindAllPatient {
	return &FindAllPatient{patientRepo: patientRepo}
}

func (u *FindAllPatient) Execute(ctx context.Context) ([]*dto.ResponsePayload, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	patients, err := u.patientRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var response []*dto.ResponsePayload
	for _, patient := range patients {
		response = append(response, dto.NewResponsePayloadPost(patient))
	}

	return response, nil
}



