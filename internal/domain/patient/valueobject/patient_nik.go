package valueobject

import "errors"

var (
	ErrorPatientNIKEmpty   = errors.New("nik tidak boleh kosong")
	ErrorPatientNIKInvalid = errors.New("nik harus 16 digit")
)

type PatientNIK struct {
	value string
}

func NewPatientNIK(value string) (PatientNIK, error) {
	// cek kondisi jika nik kosong
	if value == "" {
		return PatientNIK{}, ErrorPatientNIKEmpty
	}

	// cek kondisi jika nik tidak valid (harus 16 digit)
	if len(value) != 16 {
		return PatientNIK{}, ErrorPatientNIKInvalid
	}

	return PatientNIK{value: value}, nil
}

func ReconstructPatientNIK(value string) PatientNIK {
	return PatientNIK{value: value}
}

func (pn PatientNIK) Value() string {
	return pn.value
}

func (pn PatientNIK) Equals(other PatientNIK) bool {
	return pn.value == other.value
}
