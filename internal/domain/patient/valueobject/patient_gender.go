package valueobject

import "errors"

var (
	ErrorPatientGenderEmpty   = errors.New("jenis kelamin tidak boleh kosong")
	ErrorPatientGenderInvalid = errors.New("jenis kelamin harus laki - laki atau perempuan")
)

type PatientGender struct {
	value string
}

func NewPatientGender(value string) (PatientGender, error) {
	if value == "" {
		return PatientGender{}, ErrorPatientGenderEmpty
	}

	switch value {
	case "male", "female":
		return PatientGender{value: value}, nil
	default:
		return PatientGender{}, ErrorPatientGenderInvalid
	}
}

func ReconstructPatientGender(value string) PatientGender {
	return PatientGender{value: value}
}

func (pg PatientGender) Value() string {
	return pg.value
}

func (pg PatientGender) Equals(other PatientGender) bool {
	return pg.value == other.value
}
