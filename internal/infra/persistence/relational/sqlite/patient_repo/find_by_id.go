package patient_repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/yudhisrana/go-hospital/internal/domain/patient/entity"
)

func (r *Repository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Patient, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	query := `SELECT id, nik, name, age, gender, birth_date, address, phone, created_at, updated_at FROM patients WHERE id = ? LIMIT 1`
	row := r.db.QueryRowContext(ctx, query, id.String())

	var (
		idStr, nik, name, gender, birthDateStr, address, phone, createdAtStr, updatedAtStr string
		age                                                                                int
	)
	if err := row.Scan(&idStr, &nik, &name, &age, &gender, &birthDateStr, &address, &phone, &createdAtStr, &updatedAtStr); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error saat mencari data pasien: %w", err)
	}

	parsedID, err := uuid.Parse(idStr)
	if err != nil {
		return nil, fmt.Errorf("error saat parse ID pasien: %w", err)
	}
	birthDate, err := time.Parse(time.RFC3339, birthDateStr)
	if err != nil {
		return nil, fmt.Errorf("error saat parse tanggal lahir pasien: %w", err)
	}
	createdAt, err := time.Parse(time.RFC3339, createdAtStr)
	if err != nil {
		return nil, fmt.Errorf("error saat parse created_at pasien: %w", err)
	}
	updatedAt, err := time.Parse(time.RFC3339, updatedAtStr)
	if err != nil {
		return nil, fmt.Errorf("error saat parse updated_at pasien: %w", err)
	}

	p := entity.ReconstructPatient(parsedID, nik, name, age, gender, birthDate, address, phone, createdAt, updatedAt)
	return p, nil
}


