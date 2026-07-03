package visit_repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/yudhisrana/go-hospital/internal/domain/visit/entity"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(ctx context.Context, visit *entity.Visit) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	prescriptionJSON, _ := json.Marshal(visit.Prescriptions())

	query := `INSERT INTO visits (id, patient_id, status, queue_number, symptoms, diagnosis, prescription, registered_at, examined_at, dispensed_at, created_at, updated_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query,
		visit.ID().String(),
		visit.PatientID().String(),
		visit.Status(),
		visit.QueueNumber(),
		visit.Symptoms(),
		visit.Diagnosis(),
		string(prescriptionJSON),
		visit.RegisteredAt().Format(time.RFC3339),
		nil,
		nil,
		visit.CreatedAt().Format(time.RFC3339),
		visit.UpdatedAt().Format(time.RFC3339),
	)

	if err != nil {
		return fmt.Errorf("error saat menyimpan visit: %w", err)
	}

	return nil
}

func (r *Repository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Visit, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	query := `SELECT id, patient_id, status, queue_number, symptoms, diagnosis, prescription, registered_at, examined_at, dispensed_at, created_at, updated_at 
	FROM visits WHERE id = ? LIMIT 1`
	row := r.db.QueryRowContext(ctx, query, id.String())

	var (
		visitID, patientID, status, symptoms, diagnosis, prescriptionJSON string
		queueNumber                                                       int
		registeredAtStr, createdAtStr, updatedAtStr                       string
		examinedAtStr, dispensedAtStr                                     sql.NullString
	)

	if err := row.Scan(&visitID, &patientID, &status, &queueNumber, &symptoms, &diagnosis, &prescriptionJSON, &registeredAtStr, &examinedAtStr, &dispensedAtStr, &createdAtStr, &updatedAtStr); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error saat mencari visit: %w", err)
	}

	parsedID, _ := uuid.Parse(visitID)
	parsedPatientID, _ := uuid.Parse(patientID)
	registeredAt, _ := time.Parse(time.RFC3339, registeredAtStr)
	createdAt, _ := time.Parse(time.RFC3339, createdAtStr)
	updatedAt, _ := time.Parse(time.RFC3339, updatedAtStr)

	var examinedAt, dispensedAt *time.Time
	if examinedAtStr.Valid {
		t, _ := time.Parse(time.RFC3339, examinedAtStr.String)
		examinedAt = &t
	}
	if dispensedAtStr.Valid {
		t, _ := time.Parse(time.RFC3339, dispensedAtStr.String)
		dispensedAt = &t
	}

	var prescriptions []entity.Prescription
	if prescriptionJSON != "" {
		json.Unmarshal([]byte(prescriptionJSON), &prescriptions)
	}

	visit := entity.ReconstructVisit(parsedID, parsedPatientID, status, queueNumber, symptoms, diagnosis, prescriptions, registeredAt, examinedAt, dispensedAt, createdAt, updatedAt)
	return visit, nil
}

func (r *Repository) FindAll(ctx context.Context) ([]*entity.Visit, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	query := `SELECT id, patient_id, status, queue_number, symptoms, diagnosis, prescription, registered_at, examined_at, dispensed_at, created_at, updated_at 
	FROM visits ORDER BY queue_number ASC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error saat mengambil visits: %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()

	var result []*entity.Visit
	for rows.Next() {
		var (
			visitID, patientID, status, symptoms, diagnosis, prescriptionJSON string
			queueNumber                                                       int
			registeredAtStr, createdAtStr, updatedAtStr                       string
			examinedAtStr, dispensedAtStr                                     sql.NullString
		)

		if err := rows.Scan(&visitID, &patientID, &status, &queueNumber, &symptoms, &diagnosis, &prescriptionJSON, &registeredAtStr, &examinedAtStr, &dispensedAtStr, &createdAtStr, &updatedAtStr); err != nil {
			return nil, fmt.Errorf("error saat scan visit: %w", err)
		}

		parsedID, _ := uuid.Parse(visitID)
		parsedPatientID, _ := uuid.Parse(patientID)
		registeredAt, _ := time.Parse(time.RFC3339, registeredAtStr)
		createdAt, _ := time.Parse(time.RFC3339, createdAtStr)
		updatedAt, _ := time.Parse(time.RFC3339, updatedAtStr)

		var examinedAt, dispensedAt *time.Time
		if examinedAtStr.Valid {
			t, _ := time.Parse(time.RFC3339, examinedAtStr.String)
			examinedAt = &t
		}
		if dispensedAtStr.Valid {
			t, _ := time.Parse(time.RFC3339, dispensedAtStr.String)
			dispensedAt = &t
		}

		var prescriptions []entity.Prescription
		if prescriptionJSON != "" {
			json.Unmarshal([]byte(prescriptionJSON), &prescriptions)
		}

		visit := entity.ReconstructVisit(parsedID, parsedPatientID, status, queueNumber, symptoms, diagnosis, prescriptions, registeredAt, examinedAt, dispensedAt, createdAt, updatedAt)
		result = append(result, visit)
	}

	return result, nil
}

func (r *Repository) FindByStatus(ctx context.Context, status string) ([]*entity.Visit, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	query := `SELECT id, patient_id, status, queue_number, symptoms, diagnosis, prescription, registered_at, examined_at, dispensed_at, created_at, updated_at 
	FROM visits WHERE status = ? ORDER BY queue_number ASC`
	rows, err := r.db.QueryContext(ctx, query, status)
	if err != nil {
		return nil, fmt.Errorf("error saat mengambil visits by status: %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()

	var result []*entity.Visit
	for rows.Next() {
		var (
			visitID, patientID, st, symptoms, diagnosis, prescriptionJSON string
			queueNumber                                                   int
			registeredAtStr, createdAtStr, updatedAtStr                   string
			examinedAtStr, dispensedAtStr                                 sql.NullString
		)

		if err := rows.Scan(&visitID, &patientID, &st, &queueNumber, &symptoms, &diagnosis, &prescriptionJSON, &registeredAtStr, &examinedAtStr, &dispensedAtStr, &createdAtStr, &updatedAtStr); err != nil {
			return nil, fmt.Errorf("error saat scan visit: %w", err)
		}

		parsedID, _ := uuid.Parse(visitID)
		parsedPatientID, _ := uuid.Parse(patientID)
		registeredAt, _ := time.Parse(time.RFC3339, registeredAtStr)
		createdAt, _ := time.Parse(time.RFC3339, createdAtStr)
		updatedAt, _ := time.Parse(time.RFC3339, updatedAtStr)

		var examinedAt, dispensedAt *time.Time
		if examinedAtStr.Valid {
			t, _ := time.Parse(time.RFC3339, examinedAtStr.String)
			examinedAt = &t
		}
		if dispensedAtStr.Valid {
			t, _ := time.Parse(time.RFC3339, dispensedAtStr.String)
			dispensedAt = &t
		}

		var prescriptions []entity.Prescription
		if prescriptionJSON != "" {
			json.Unmarshal([]byte(prescriptionJSON), &prescriptions)
		}

		visit := entity.ReconstructVisit(parsedID, parsedPatientID, st, queueNumber, symptoms, diagnosis, prescriptions, registeredAt, examinedAt, dispensedAt, createdAt, updatedAt)
		result = append(result, visit)
	}

	return result, nil
}

func (r *Repository) Update(ctx context.Context, visit *entity.Visit) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	prescriptionJSON, _ := json.Marshal(visit.Prescriptions())

	query := `UPDATE visits SET status = ?, diagnosis = ?, prescription = ?, examined_at = ?, dispensed_at = ?, updated_at = ? WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query,
		visit.Status(),
		visit.Diagnosis(),
		string(prescriptionJSON),
		visit.ExaminedAt(),
		visit.DispensedAt(),
		visit.UpdatedAt().Format(time.RFC3339),
		visit.ID().String(),
	)

	if err != nil {
		return fmt.Errorf("error saat update visit: %w", err)
	}

	return nil
}

func (r *Repository) GetNextQueueNumber(ctx context.Context) (int, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}

	query := `SELECT COALESCE(MAX(queue_number), 0) + 1 FROM visits WHERE registered_at >= date('now')`
	var nextNumber int

	if err := r.db.QueryRowContext(ctx, query).Scan(&nextNumber); err != nil {
		return 0, fmt.Errorf("error saat get next queue number: %w", err)
	}

	return nextNumber, nil
}
