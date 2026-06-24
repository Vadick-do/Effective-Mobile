package subscriptions_service

import (
	"context"
	"fmt"

	"github.com/Vadick-do/Effective-Mobile/internal/core/domain"
	"github.com/google/uuid"
)

func (s *SubscriptionsService) PutSubscription(
	ctx context.Context,
	subID uuid.UUID,
	subscription domain.Subscription,
) (domain.Subscription, error) {
	subscription, err := s.subscriptionsRepository.PutSubscription(ctx, subID, subscription)
	if err != nil {
		return domain.Subscription{}, fmt.Errorf("put subscsription: %w", err)
	}

	return subscription, nil
}
