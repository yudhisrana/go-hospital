package patient_repo

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/yudhisrana/go-hospital/internal/domain/patient/entity"
)

func (r *Repository) FindAll(ctx context.Context) ([]*entity.Patient, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	query := `SELECT id, nik, name, age, gender, birth_date, address, phone, created_at, updated_at FROM patients`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error saat mengambil data pasien: %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()

	var result []*entity.Patient
	for rows.Next() {
		var (
			idStr, nik, name, gender, birthDateStr, address, phone, createdAtStr, updatedAtStr string
			age                                                                               int
		)
		if err := rows.Scan(&idStr, &nik, &name, &age, &gender, &birthDateStr, &address, &phone, &createdAtStr, &updatedAtStr); err != nil {
			return nil, fmt.Errorf("error saat scan data pasien: %w", err)
		}

		id, err := uuid.Parse(idStr)
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

		p := entity.ReconstructPatient(id, nik, name, age, gender, birthDate, address, phone, createdAt, updatedAt)
		result = append(result, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error saat iterasi data pasien: %w", err)
	}

	return result, nil
}



