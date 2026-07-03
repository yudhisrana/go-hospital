package patient_repo

import (
	"context"
	"fmt"
	"time"

	"github.com/yudhisrana/go-hospital/internal/domain/patient/entity"
)

func (r *Repository) Update(ctx context.Context, patient *entity.Patient) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	query := `UPDATE patients SET nik = ?, name = ?, age = ?, gender = ?, birth_date = ?, address = ?, phone = ?, updated_at = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query,
		patient.NIK().Value(),
		patient.Name().Value(),
		patient.Age().Value(),
		patient.Gender().Value(),
		patient.BirthDate().Value().Format(time.RFC3339),
		patient.Address().Value(),
		patient.Phone().Value(),
		patient.UpdatedAt().Format(time.RFC3339),
		patient.ID().String(),
	)

	if err != nil {
		return fmt.Errorf("error saat memperbarui data pasien: %w", err)
	}

	return nil
}


