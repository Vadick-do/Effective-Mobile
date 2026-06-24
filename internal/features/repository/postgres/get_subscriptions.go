package subscriptions_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Vadick-do/Effective-Mobile/internal/core/domain"
	core_errors "github.com/Vadick-do/Effective-Mobile/internal/core/errors"
	core_postgres_pool "github.com/Vadick-do/Effective-Mobile/internal/core/repository/postgres/pool"
	"github.com/google/uuid"
)

func (r *SubscriptionsRepository) GetSubscription(
	ctx context.Context,
	subID uuid.UUID,
) (domain.Subscription, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, service_name, price, user_id, start_date, end_date, created_at, updated_at
	FROM efmobapp.subscriptions
	WHERE id=$1;
	`

	row := r.pool.QueryRow(ctx, query, subID)

	var subModel SubscriptionsModel
	err := row.Scan(
		&subModel.ID,
		&subModel.ServiceName,
		&subModel.Price,
		&subModel.UserID,
		&subModel.StartDate,
		&subModel.EndDate,
		&subModel.CreatedAt,
		&subModel.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Subscription{}, fmt.Errorf(
				"subscription with id=%v: %w",
				subID,
				core_errors.ErrNotFound,
			)
		}
		return domain.Subscription{}, fmt.Errorf("scan error: %w", err)
	}

	subDomain := subscriptionDomainFromModel(subModel)
	return subDomain, nil
}
