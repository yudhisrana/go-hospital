package valueobject

import "errors"

var (
	ErrorPatientAgeInvalid = errors.New("umur harus lebih dari 0")
)

type PatientAge struct {
	value int
}

func NewPatientAge(value int) (PatientAge, error) {
	if value <= 0 {
		return PatientAge{}, ErrorPatientAgeInvalid
	}

	return PatientAge{value: value}, nil
}

func ReconstructPatientAge(value int) PatientAge {
	return PatientAge{value: value}
}

func (pa PatientAge) Value() int {
	return pa.value
}

func (pa PatientAge) Equals(other PatientAge) bool {
	return pa.value == other.value
}

