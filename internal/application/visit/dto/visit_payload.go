package dto

import (
	"github.com/yudhisrana/go-hospital/internal/domain/visit/entity"
)

type RegisterVisitRequest struct {
	PatientID string `json:"patient_id"`
	Symptoms  string `json:"symptoms"`
}

type ExaminePatientRequest struct {
	Diagnosis     string              `json:"diagnosis"`
	Prescriptions []PrescriptionInput `json:"prescriptions"`
}

type PrescriptionInput struct {
	Medicine string `json:"medicine"`
	Dosage   string `json:"dosage"`
	Quantity int    `json:"quantity"`
}

type VisitResponse struct {
	ID            string               `json:"id"`
	PatientID     string               `json:"patient_id"`
	Status        string               `json:"status"`
	QueueNumber   int                  `json:"queue_number"`
	Symptoms      string               `json:"symptoms"`
	Diagnosis     string               `json:"diagnosis,omitempty"`
	Prescriptions []PrescriptionOutput `json:"prescriptions,omitempty"`
	RegisteredAt  string               `json:"registered_at"`
	ExaminedAt    *string              `json:"examined_at,omitempty"`
	DispensedAt   *string              `json:"dispensed_at,omitempty"`
	CreatedAt     string               `json:"created_at"`
	UpdatedAt     string               `json:"updated_at"`
}

type PrescriptionOutput struct {
	Medicine string `json:"medicine"`
	Dosage   string `json:"dosage"`
	Quantity int    `json:"quantity"`
}

func NewVisitResponse(visit *entity.Visit) *VisitResponse {
	var examinedAt, dispensedAt *string
	if visit.ExaminedAt() != nil {
		t := visit.ExaminedAt().Format("2006-01-02 15:04:05")
		examinedAt = &t
	}
	if visit.DispensedAt() != nil {
		t := visit.DispensedAt().Format("2006-01-02 15:04:05")
		dispensedAt = &t
	}

	prescriptions := make([]PrescriptionOutput, len(visit.Prescriptions()))
	for i, p := range visit.Prescriptions() {
		prescriptions[i] = PrescriptionOutput{
			Medicine: p.Medicine,
			Dosage:   p.Dosage,
			Quantity: p.Quantity,
		}
	}

	return &VisitResponse{
		ID:            visit.ID().String(),
		PatientID:     visit.PatientID().String(),
		Status:        visit.Status(),
		QueueNumber:   visit.QueueNumber(),
		Symptoms:      visit.Symptoms(),
		Diagnosis:     visit.Diagnosis(),
		Prescriptions: prescriptions,
		RegisteredAt:  visit.RegisteredAt().Format("2006-01-02 15:04:05"),
		ExaminedAt:    examinedAt,
		DispensedAt:   dispensedAt,
		CreatedAt:     visit.CreatedAt().Format("2006-01-02 15:04:05"),
		UpdatedAt:     visit.UpdatedAt().Format("2006-01-02 15:04:05"),
	}
}
