package subscriptions_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/Vadick-do/Effective-Mobile/internal/core/errors"
	"github.com/google/uuid"
)

func (r *SubscriptionsRepository) DeleteSubcsription(
	ctx context.Context,
	subID uuid.UUID,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	DELETE FROM efmobapp.subscriptions
	WHERE id=$1;
	`

	cmdTag, err := r.pool.Exec(ctx, query, subID)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("subscription with id=%v: %w", subID, core_errors.ErrNotFound)
	}

	return nil
}
