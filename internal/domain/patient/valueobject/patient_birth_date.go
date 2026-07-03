package valueobject

import (
	"errors"
	"time"
)

var (
	ErrorPatientBirthDateZero = errors.New("tanggal lahir tidak boleh kosong")
	ErrorPatientBirthDateFuture = errors.New("tanggal lahir tidak boleh di masa depan")
)

type PatientBirthDate struct {
	value time.Time
}

func NewPatientBirthDate(value time.Time) (PatientBirthDate, error) {
	if value.IsZero() {
		return PatientBirthDate{}, ErrorPatientBirthDateZero
	}

	if value.After(time.Now()) {
		return PatientBirthDate{}, ErrorPatientBirthDateFuture
	}

	return PatientBirthDate{value: value}, nil
}

func ReconstructPatientBirthDate(value time.Time) PatientBirthDate {
	return PatientBirthDate{value: value}
}

func (pbd PatientBirthDate) Value() time.Time {
	return pbd.value
}

func (pbd PatientBirthDate) Equals(other PatientBirthDate) bool {
	return pbd.value.Equal(other.value)
}

