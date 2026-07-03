package valueobject

import "errors"

var (
	ErrorPatientAddressEmpty = errors.New("alamat tidak boleh kosong")
)

type PatientAddress struct {
	value string
}

func NewPatientAddress(value string) (PatientAddress, error) {
	if value == "" {
		return PatientAddress{}, ErrorPatientAddressEmpty
	}

	return PatientAddress{value: value}, nil
}

func ReconstructPatientAddress(value string) PatientAddress {
	return PatientAddress{value: value}
}

func (pa PatientAddress) Value() string {
	return pa.value
}

func (pa PatientAddress) Equals(other PatientAddress) bool {
	return pa.value == other.value
}
