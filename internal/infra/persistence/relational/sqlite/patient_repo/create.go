package patient_repo

import (
	"context"
	"fmt"
	"time"

	"github.com/yudhisrana/go-hospital/internal/domain/patient/entity"
)

func (r *Repository) Save(ctx context.Context, patient *entity.Patient) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	query := `INSERT INTO patients (id, nik, name, age, gender, birth_date, address, phone, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query,
		patient.ID().String(),
		patient.NIK().Value(),
		patient.Name().Value(),
		patient.Age().Value(),
		patient.Gender().Value(),
		patient.BirthDate().Value().Format(time.RFC3339),
		patient.Address().Value(),
		patient.Phone().Value(),
		patient.CreatedAt().Format(time.RFC3339),
		patient.UpdatedAt().Format(time.RFC3339),
	)

	if err != nil {
		return fmt.Errorf("error saat menyimpan data pasien: %w", err)
	}

	return nil
}
