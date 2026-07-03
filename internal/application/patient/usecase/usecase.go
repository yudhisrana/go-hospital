package usecase

import "github.com/yudhisrana/go-hospital/internal/domain/patient/repository"

type PatientRepository interface {
	repository.CreatePatientRepository
	repository.UpdatePatientRepository
	repository.DeletePatientRepository
	repository.FindPatientByIDRepository
	repository.FindAllPatientRepository
}

type PatientUseCase struct {
	CreatePatient   *CreatePatient
	UpdatePatient   *UpdatePatient
	DeletePatient   *DeletePatient
	FindPatientByID *FindPatientByID
	FindAllPatient  *FindAllPatient
}

func NewPatientUseCase(repo PatientRepository) *PatientUseCase {
	return &PatientUseCase{
		FindAllPatient: NewFindAllPatientUseCase(PatientFindAllUsecase{
			FindAllPatientRepository: repo,
		}),
		FindPatientByID: NewFindPatientByIDUseCase(PatientFindByIDUsecase{
			FindPatientByIDRepository: repo,
		}),
		CreatePatient: NewCreatePatientUseCase(PatientCreateUsecase{
			CreatePatientRepository: repo,
		}),
		UpdatePatient: NewUpdatePatientUseCase(PatientUpdateUsecase{
			UpdatePatientRepository:   repo,
			FindPatientByIDRepository: repo,
		}),
		DeletePatient: NewDeletePatientUseCase(DeletePatientUseCase{
			FindPatientByIDRepository: repo,
			DeletePatientRepository:   repo,
		}),
	}
}


