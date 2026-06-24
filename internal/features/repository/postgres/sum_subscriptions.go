package subscriptions_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Vadick-do/Effective-Mobile/internal/core/domain"
	"github.com/google/uuid"
)

func (r *SubscriptionsRepository) SumSubscriptions(
	ctx context.Context,
	from string,
	to string,
	userID *uuid.UUID,
	serviceName *string,
) (domain.Total, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT COALESCE(SUM(price), 0) FROM efmobapp.subscriptions WHERE start_date <= $1 AND (end_date IS NULL OR end_date >= $2)`
	args := []interface{}{to, from}

	if userID != nil {
		query += fmt.Sprintf(" AND user_id = $%d", len(args)+1)
		args = append(args, *userID)
	}

	if serviceName != nil {
		query += fmt.Sprintf(" AND service_name = $%d", len(args)+1)
		args = append(args, *serviceName)
	}

	row := r.pool.QueryRow(ctx, query, args...)

	var total int
	if err := row.Scan(&total); err != nil {
		return domain.Total{}, fmt.Errorf("scan total price: %w", err)
	}

	return domain.Total{Total: total}, nil
}
