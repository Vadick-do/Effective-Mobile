package subscriptions_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Vadick-do/Effective-Mobile/internal/core/domain"
	core_errors "github.com/Vadick-do/Effective-Mobile/internal/core/errors"
	core_postgres_pool "github.com/Vadick-do/Effective-Mobile/internal/core/repository/postgres/pool"
	"github.com/google/uuid"
)

func (r *SubscriptionsRepository) PutSubscription(
	ctx context.Context,
	subID uuid.UUID,
	subscription domain.Subscription,
) (domain.Subscription, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	upTime := time.Now().UTC()

	query := `
	UPDATE efmobapp.subscriptions 
	SET
		service_name=$1,
		price=$2, 
		user_id=$3, 
		start_date=$4, 
		end_date=$5,
		updated_at=$6
	WHERE id=$7

	RETURNING 
		id,
		service_name,
		price,
		user_id,
		start_date,
		end_date,
		created_at,
		updated_at;
	`

	row := r.pool.QueryRow(ctx, query, subscription.ServiceName, subscription.Price, subscription.UserID, subscription.StartDate, subscription.EndDate, upTime, subID)

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
			return domain.Subscription{}, fmt.Errorf("subscription with id='%v' concurrently accessed: %w", subID, core_errors.ErrConflict)
		}
		return domain.Subscription{}, fmt.Errorf("scan error: %w", err)
	}

	subscriptionDomain := subscriptionDomainFromModel(subModel)
	return subscriptionDomain, nil
}
