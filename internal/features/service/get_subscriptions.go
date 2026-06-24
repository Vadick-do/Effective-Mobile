package subscriptions_service

import (
	"context"
	"fmt"

	"github.com/Vadick-do/Effective-Mobile/internal/core/domain"
	"github.com/google/uuid"
)

func (s *SubscriptionsService) GetSubscription(
	ctx context.Context,
	subID uuid.UUID,
) (domain.Subscription, error) {
	subscription, err := s.subscriptionsRepository.GetSubscription(ctx, subID)
	if err != nil {
		return domain.Subscription{}, fmt.Errorf("get subscription: %w", err)
	}

	return subscription, nil
}
