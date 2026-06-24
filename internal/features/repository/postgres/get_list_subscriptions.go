package subscriptions_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Vadick-do/Effective-Mobile/internal/core/domain"
)

func (r *SubscriptionsRepository) GetSubscriptions(
	ctx context.Context,
) ([]domain.Subscription, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, service_name, price, user_id, start_date, end_date, created_at, updated_at
	FROM efmobapp.subscriptions
	ORDER BY created_at DESC;
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("select subscriptions: %w", err)
	}
	defer rows.Close()

	var subsModels []SubscriptionsModel
	for rows.Next() {
		var subModel SubscriptionsModel

		err := rows.Scan(
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
			return nil, fmt.Errorf("scan subscriptions: %w", err)
		}

		subsModels = append(subsModels, subModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	subsDomains := subscriptionDomainsFromModels(subsModels)
	return subsDomains, nil
}
