package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/yudhisrana/go-hospital/internal/domain/visit/entity"
	"github.com/yudhisrana/go-hospital/internal/domain/visit/repository"
)

type VisitRepository interface {
	repository.CreateVisitRepository
	repository.FindVisitRepository
	repository.UpdateVisitRepository
	repository.GetNextQueueNumberRepository
}

type VisitUseCase struct {
	RegisterVisit     *RegisterVisit
	ExaminePatient    *ExaminePatient
	DispenseMedicine  *DispenseMedicine
	FindAllVisit      *FindAllVisit
	FindVisitByID     *FindVisitByID
	FindVisitByStatus *FindVisitByStatus
}

func NewVisitUseCase(repo VisitRepository) *VisitUseCase {
	return &VisitUseCase{
		RegisterVisit:     NewRegisterVisitUseCase(repo),
		ExaminePatient:    NewExaminePatientUseCase(repo),
		DispenseMedicine:  NewDispenseMedicineUseCase(repo),
		FindAllVisit:      NewFindAllVisitUseCase(repo),
		FindVisitByID:     NewFindVisitByIDUseCase(repo),
		FindVisitByStatus: NewFindVisitByStatusUseCase(repo),
	}
}

// RegisterVisit usecase
type RegisterVisit struct {
	repo VisitRepository
}

func NewRegisterVisitUseCase(repo VisitRepository) *RegisterVisit {
	return &RegisterVisit{repo: repo}
}

func (u *RegisterVisit) Execute(ctx context.Context, patientID uuid.UUID, symptoms string) (*entity.Visit, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	queueNumber, err := u.repo.GetNextQueueNumber(ctx)
	if err != nil {
		return nil, err
	}

	visit, err := entity.NewVisit(patientID, queueNumber, symptoms)
	if err != nil {
		return nil, err
	}

	if err := u.repo.Save(ctx, visit); err != nil {
		return nil, err
	}

	return visit, nil
}

// ExaminePatient usecase
type ExaminePatient struct {
	repo VisitRepository
}

func NewExaminePatientUseCase(repo VisitRepository) *ExaminePatient {
	return &ExaminePatient{repo: repo}
}

func (u *ExaminePatient) Execute(ctx context.Context, visitID uuid.UUID, diagnosis string, prescriptions []entity.Prescription) (*entity.Visit, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	visit, err := u.repo.FindByID(ctx, visitID)
	if err != nil {
		return nil, err
	}
	if visit == nil {
		return nil, entity.ErrVisitNotFound
	}

	if err := visit.ExaminePatient(diagnosis, prescriptions); err != nil {
		return nil, err
	}

	if err := u.repo.Update(ctx, visit); err != nil {
		return nil, err
	}

	return visit, nil
}

// DispenseMedicine usecase
type DispenseMedicine struct {
	repo VisitRepository
}

func NewDispenseMedicineUseCase(repo VisitRepository) *DispenseMedicine {
	return &DispenseMedicine{repo: repo}
}

func (u *DispenseMedicine) Execute(ctx context.Context, visitID uuid.UUID) (*entity.Visit, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	visit, err := u.repo.FindByID(ctx, visitID)
	if err != nil {
		return nil, err
	}
	if visit == nil {
		return nil, entity.ErrVisitNotFound
	}

	if err := visit.DispenseMedicine(); err != nil {
		return nil, err
	}

	if err := u.repo.Update(ctx, visit); err != nil {
		return nil, err
	}

	return visit, nil
}

// FindAllVisit usecase
type FindAllVisit struct {
	repo VisitRepository
}

func NewFindAllVisitUseCase(repo VisitRepository) *FindAllVisit {
	return &FindAllVisit{repo: repo}
}

func (u *FindAllVisit) Execute(ctx context.Context) ([]*entity.Visit, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return u.repo.FindAll(ctx)
}

// FindVisitByID usecase
type FindVisitByID struct {
	repo VisitRepository
}

func NewFindVisitByIDUseCase(repo VisitRepository) *FindVisitByID {
	return &FindVisitByID{repo: repo}
}

func (u *FindVisitByID) Execute(ctx context.Context, visitID uuid.UUID) (*entity.Visit, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return u.repo.FindByID(ctx, visitID)
}

// FindVisitByStatus usecase
type FindVisitByStatus struct {
	repo VisitRepository
}

func NewFindVisitByStatusUseCase(repo VisitRepository) *FindVisitByStatus {
	return &FindVisitByStatus{repo: repo}
}

func (u *FindVisitByStatus) Execute(ctx context.Context, status string) ([]*entity.Visit, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return u.repo.FindByStatus(ctx, status)
}
