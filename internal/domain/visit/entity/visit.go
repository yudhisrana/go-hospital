package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	VisitStatusRegistered      = "registered"
	VisitStatusWaitingDoctor   = "waiting_doctor"
	VisitStatusWaitingPharmacy = "waiting_pharmacy"
	VisitStatusCompleted       = "completed"
)

var (
	ErrInvalidVisitStatus = errors.New("invalid visit status")
	ErrVisitNotFound      = errors.New("visit not found")
	ErrInvalidSymptoms    = errors.New("symptoms cannot be empty")
)

type Prescription struct {
	Medicine string `json:"medicine"`
	Dosage   string `json:"dosage"`
	Quantity int    `json:"quantity"`
}

type Visit struct {
	id            uuid.UUID
	patientID     uuid.UUID
	status        string
	queueNumber   int
	symptoms      string
	diagnosis     string
	prescriptions []Prescription
	registeredAt  time.Time
	examinedAt    *time.Time
	dispensedAt   *time.Time
	createdAt     time.Time
	updatedAt     time.Time
}

func NewVisit(patientID uuid.UUID, queueNumber int, symptoms string) (*Visit, error) {
	if symptoms == "" {
		return nil, ErrInvalidSymptoms
	}

	return &Visit{
		id:           uuid.New(),
		patientID:    patientID,
		status:       VisitStatusRegistered,
		queueNumber:  queueNumber,
		symptoms:     symptoms,
		registeredAt: time.Now(),
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}, nil
}

func ReconstructVisit(id, patientID uuid.UUID, status string, queueNumber int, symptoms, diagnosis string, prescriptions []Prescription, registeredAt time.Time, examinedAt, dispensedAt *time.Time, createdAt, updatedAt time.Time) *Visit {
	return &Visit{
		id:            id,
		patientID:     patientID,
		status:        status,
		queueNumber:   queueNumber,
		symptoms:      symptoms,
		diagnosis:     diagnosis,
		prescriptions: prescriptions,
		registeredAt:  registeredAt,
		examinedAt:    examinedAt,
		dispensedAt:   dispensedAt,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
	}
}

func (v *Visit) ExaminePatient(diagnosis string, prescriptions []Prescription) error {
	if diagnosis == "" {
		return errors.New("diagnosis cannot be empty")
	}
	if len(prescriptions) == 0 {
		return errors.New("at least one prescription required")
	}

	v.diagnosis = diagnosis
	v.prescriptions = prescriptions
	v.status = VisitStatusWaitingPharmacy
	now := time.Now()
	v.examinedAt = &now
	v.updatedAt = time.Now()

	return nil
}

func (v *Visit) DispenseMedicine() error {
	if v.status != VisitStatusWaitingPharmacy {
		return errors.New("visit must be in waiting_pharmacy status")
	}

	v.status = VisitStatusCompleted
	now := time.Now()
	v.dispensedAt = &now
	v.updatedAt = time.Now()

	return nil
}

func (v *Visit) ID() uuid.UUID {
	return v.id
}

func (v *Visit) PatientID() uuid.UUID {
	return v.patientID
}

func (v *Visit) Status() string {
	return v.status
}

func (v *Visit) QueueNumber() int {
	return v.queueNumber
}

func (v *Visit) Symptoms() string {
	return v.symptoms
}

func (v *Visit) Diagnosis() string {
	return v.diagnosis
}

func (v *Visit) Prescriptions() []Prescription {
	return v.prescriptions
}

func (v *Visit) RegisteredAt() time.Time {
	return v.registeredAt
}

func (v *Visit) ExaminedAt() *time.Time {
	return v.examinedAt
}

func (v *Visit) DispensedAt() *time.Time {
	return v.dispensedAt
}

func (v *Visit) CreatedAt() time.Time {
	return v.createdAt
}

func (v *Visit) UpdatedAt() time.Time {
	return v.updatedAt
}
