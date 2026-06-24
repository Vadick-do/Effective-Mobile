package subscriptions_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Vadick-do/Effective-Mobile/internal/core/domain"
	core_errors "github.com/Vadick-do/Effective-Mobile/internal/core/errors"
	core_postgres_pool "github.com/Vadick-do/Effective-Mobile/internal/core/repository/postgres/pool"
)

func (r *SubscriptionsRepository) CreateSubscription(
	ctx context.Context,
	subscription domain.Subscription,
) (domain.Subscription, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `INSERT INTO efmobapp.subscriptions (id, service_name, price, user_id, start_date, end_date, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id, service_name, price, user_id, start_date, end_date, created_at, updated_at;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		subscription.ID,
		subscription.ServiceName,
		subscription.Price,
		subscription.UserID,
		subscription.StartDate,
		subscription.EndDate,
		subscription.CreatedAt,
		subscription.UpdatedAt,
	)

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
		// На случай добавления внешнего ключа
		if errors.Is(err, core_postgres_pool.ErrViolatesForeignKey) {
			return domain.Subscription{}, fmt.Errorf("%v: subscription with user_id=`%v`: %w", err, subscription.UserID, core_errors.ErrNotFound)
		}
		return domain.Subscription{}, fmt.Errorf("scan error: %w", err)
	}

	subscriptionDomain := subscriptionDomainFromModel(subModel)

	return subscriptionDomain, nil
}
