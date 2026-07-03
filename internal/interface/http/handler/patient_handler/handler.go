package patient_handler

import "github.com/yudhisrana/go-hospital/internal/application/patient/usecase"

type Handler struct {
	usc *usecase.PatientUseCase
}

func NewPatientHandler(usc *usecase.PatientUseCase) *Handler {
	return &Handler{
		usc: usc,
	}
}
