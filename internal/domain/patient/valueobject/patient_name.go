package valueobject

import "errors"

var (
	ErrorPatientNameEmpty   = errors.New("nama tidak boleh kosong")
	ErrorPatientNameInvalid = errors.New("nama minimal 3 karakter")
)

type PatientName struct {
	value string
}

func NewPatientName(value string) (PatientName, error) {
	// cek kondisi jika nama kosong
	if value == "" {
		return PatientName{}, ErrorPatientNameEmpty
	}

	// cek kondisi jika nama tidak valid (harus 3 karakter)
	if len(value) < 3 {
		return PatientName{}, ErrorPatientNameInvalid
	}

	return PatientName{value: value}, nil
}

func ReconstructPatientName(value string) PatientName {
	return PatientName{value: value}
}

func (pn PatientName) Value() string {
	return pn.value
}

func (pn PatientName) Equals(other PatientName) bool {
	return pn.value == other.value
}
