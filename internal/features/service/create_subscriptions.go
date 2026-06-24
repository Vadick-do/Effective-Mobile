package subscriptions_service

import (
	"context"
	"fmt"

	"github.com/Vadick-do/Effective-Mobile/internal/core/domain"
)

func (s *SubscriptionsService) CreateSubscription(
	ctx context.Context,
	subscription domain.Subscription,
) (domain.Subscription, error) {
	subscription, err := s.subscriptionsRepository.CreateSubscription(ctx, subscription)
	if err != nil {
		return domain.Subscription{}, fmt.Errorf("create subscription: %w", err)
	}

	return subscription, nil
}
