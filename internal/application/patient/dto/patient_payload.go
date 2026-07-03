package dto

import (
	"time"

	"github.com/yudhisrana/go-hospital/internal/domain/patient/entity"
)

type RequestPayload struct {
	NIK         string    `json:"nik"`
	Name        string    `json:"name"`
	Age         int       `json:"age"`
	Gender      string    `json:"gender"`
	BirthDate   time.Time `json:"birth_date"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
}

type ResponsePayload struct {
	ID          string    `json:"id"`
	NIK         string    `json:"nik"`
	Name        string    `json:"name"`
	Age         int       `json:"age"`
	Gender      string    `json:"gender"`
	BirthDate   time.Time `json:"birth_date"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}

func NewResponsePayloadPost(patient *entity.Patient) *ResponsePayload {
	return &ResponsePayload{
		ID:          patient.ID().String(),
		NIK:         patient.NIK().Value(),
		Name:        patient.Name().Value(),
		Age:         patient.Age().Value(),
		Gender:      patient.Gender().Value(),
		BirthDate:   patient.BirthDate().Value(),
		Address:     patient.Address().Value(),
		PhoneNumber: patient.Phone().Value(),
		CreatedAt:   patient.CreatedAt().Format("2006-01-02 15:04:05"),
		UpdatedAt:   patient.UpdatedAt().Format("2006-01-02 15:04:05"),
	}
}
