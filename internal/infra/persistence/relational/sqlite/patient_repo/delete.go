package patient_repo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	query := `DELETE FROM patients WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id.String())

	if err != nil {
		return fmt.Errorf("error saat menghapus data pasien: %w", err)
	}

	return nil
}


