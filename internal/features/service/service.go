package subscriptions_service

import (
	"context"

	"github.com/Vadick-do/Effective-Mobile/internal/core/domain"
	"github.com/google/uuid"
)

type SubscriptionsService struct {
	subscriptionsRepository SubscriptionsRepository
}

type SubscriptionsRepository interface {
	CreateSubscription(
		ctx context.Context,
		subscription domain.Subscription,
	) (domain.Subscription, error)
	GetSubscription(
		ctx context.Context,
		subID uuid.UUID,
	) (domain.Subscription, error)
	DeleteSubcsription(
		ctx context.Context,
		subID uuid.UUID,
	) error
	GetSubscriptions(
		ctx context.Context,
	) ([]domain.Subscription, error)
	PutSubscription(
		ctx context.Context,
		subID uuid.UUID,
		subscription domain.Subscription,
	) (domain.Subscription, error)
	SumSubscriptions(
		ctx context.Context,
		from string,
		to string,
		userID *uuid.UUID,
		serviceName *string,
	) (domain.Total, error)
}

func NewSubscriptionsService(
	subscriptionsRepository SubscriptionsRepository,
) *SubscriptionsService {
	return &SubscriptionsService{
		subscriptionsRepository: subscriptionsRepository,
	}
}
