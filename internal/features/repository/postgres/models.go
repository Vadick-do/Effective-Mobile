package subscriptions_postgres_repository

import (
	"time"

	"github.com/Vadick-do/Effective-Mobile/internal/core/domain"
	"github.com/google/uuid"
)

type SubscriptionsModel struct {
	ID          uuid.UUID
	ServiceName string
	Price       int
	UserID      uuid.UUID
	StartDate   string
	EndDate     *string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func subscriptionDomainsFromModels(subModels []SubscriptionsModel) []domain.Subscription {
	domains := make([]domain.Subscription, len(subModels))

	for i, model := range subModels {
		domains[i] = subscriptionDomainFromModel(model)
	}

	return domains
}

func subscriptionDomainFromModel(subModel SubscriptionsModel) domain.Subscription {
	return domain.NewSubscription(
		subModel.ID,
		subModel.ServiceName,
		subModel.Price,
		subModel.UserID,
		subModel.StartDate,
		subModel.EndDate,
		subModel.CreatedAt,
		subModel.UpdatedAt,
	)
}
