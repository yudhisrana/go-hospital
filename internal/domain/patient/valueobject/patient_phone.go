package valueobject

import "errors"

var (
	ErrorPatientPhoneEmpty   = errors.New("nomor telepon tidak boleh kosong")
	ErrorPatientPhoneInvalid = errors.New("nomor telepon minimal 10 digit")
)

type PatientPhone struct {
	value string
}

func NewPatientPhone(value string) (PatientPhone, error) {
	if value == "" {
		return PatientPhone{}, ErrorPatientPhoneEmpty
	}

	if len(value) < 10 {
		return PatientPhone{}, ErrorPatientPhoneInvalid
	}

	return PatientPhone{value: value}, nil
}

func ReconstructPatientPhone(value string) PatientPhone {
	return PatientPhone{value: value}
}

func (pp PatientPhone) Value() string {
	return pp.value
}

func (pp PatientPhone) Equals(other PatientPhone) bool {
	return pp.value == other.value
}

