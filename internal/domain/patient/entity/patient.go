package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/yudhisrana/go-hospital/internal/domain/patient/valueobject"
)

var (
	ErrPatientNotChange = errors.New("data pasien tidak berubah")
)

type Patient struct {
	id        uuid.UUID
	nik       valueobject.PatientNIK
	name      valueobject.PatientName
	age       valueobject.PatientAge
	gender    valueobject.PatientGender
	birthDate valueobject.PatientBirthDate
	address   valueobject.PatientAddress
	phone     valueobject.PatientPhone
	createdAt time.Time
	updatedAt time.Time
}

func NewPatient(nik, name string, age int, gender string, birthDate time.Time, address, phone string) (*Patient, error) {
	nikVO, err := valueobject.NewPatientNIK(nik)
	if err != nil {
		return nil, err
	}

	nameVO, err := valueobject.NewPatientName(name)
	if err != nil {
		return nil, err
	}

	ageVO, err := valueobject.NewPatientAge(age)
	if err != nil {
		return nil, err
	}

	genderVO, err := valueobject.NewPatientGender(gender)
	if err != nil {
		return nil, err
	}

	birthDateVO, err := valueobject.NewPatientBirthDate(birthDate)
	if err != nil {
		return nil, err
	}

	addressVO, err := valueobject.NewPatientAddress(address)
	if err != nil {
		return nil, err
	}

	phoneVO, err := valueobject.NewPatientPhone(phone)
	if err != nil {
		return nil, err
	}

	return &Patient{
		id:        uuid.New(),
		nik:       nikVO,
		name:      nameVO,
		age:       ageVO,
		gender:    genderVO,
		birthDate: birthDateVO,
		address:   addressVO,
		phone:     phoneVO,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func ReconstructPatient(id uuid.UUID, nik, name string, age int, gender string, birthDate time.Time, address, phone string, createdAt, updatedAt time.Time) *Patient {
	nikVO := valueobject.ReconstructPatientNIK(nik)
	nameVO := valueobject.ReconstructPatientName(name)
	ageVO := valueobject.ReconstructPatientAge(age)
	genderVO := valueobject.ReconstructPatientGender(gender)
	birthDateVO := valueobject.ReconstructPatientBirthDate(birthDate)
	addressVO := valueobject.ReconstructPatientAddress(address)
	phoneVO := valueobject.ReconstructPatientPhone(phone)

	return &Patient{
		id:        id,
		nik:       nikVO,
		name:      nameVO,
		age:       ageVO,
		gender:    genderVO,
		birthDate: birthDateVO,
		address:   addressVO,
		phone:     phoneVO,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (p *Patient) UpdatePatient(nik, name string, age int, gender string, birthDate time.Time, address, phone string) error {
	nikVO, err := valueobject.NewPatientNIK(nik)
	if err != nil {
		return err
	}

	nameVO, err := valueobject.NewPatientName(name)
	if err != nil {
		return err
	}

	ageVO, err := valueobject.NewPatientAge(age)
	if err != nil {
		return err
	}

	genderVO, err := valueobject.NewPatientGender(gender)
	if err != nil {
		return err
	}

	birthDateVO, err := valueobject.NewPatientBirthDate(birthDate)
	if err != nil {
		return err
	}

	addressVO, err := valueobject.NewPatientAddress(address)
	if err != nil {
		return err
	}

	phoneVO, err := valueobject.NewPatientPhone(phone)
	if err != nil {
		return err
	}

	if p.nik.Equals(nikVO) && p.name.Equals(nameVO) && p.age.Equals(ageVO) && p.gender.Equals(genderVO) && p.birthDate.Equals(birthDateVO) && p.address.Equals(addressVO) && p.phone.Equals(phoneVO) {
		return ErrPatientNotChange
	}

	p.nik = nikVO
	p.name = nameVO
	p.age = ageVO
	p.gender = genderVO
	p.birthDate = birthDateVO
	p.address = addressVO
	p.phone = phoneVO
	p.updatedAt = time.Now()

	return nil
}

func (p *Patient) ID() uuid.UUID {
	return p.id
}

func (p *Patient) NIK() valueobject.PatientNIK {
	return p.nik
}

func (p *Patient) Name() valueobject.PatientName {
	return p.name
}

func (p *Patient) Age() valueobject.PatientAge {
	return p.age
}

func (p *Patient) Gender() valueobject.PatientGender {
	return p.gender
}

func (p *Patient) BirthDate() valueobject.PatientBirthDate {
	return p.birthDate
}

func (p *Patient) Address() valueobject.PatientAddress {
	return p.address
}

func (p *Patient) Phone() valueobject.PatientPhone {
	return p.phone
}

func (p *Patient) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Patient) UpdatedAt() time.Time {
	return p.updatedAt
}
