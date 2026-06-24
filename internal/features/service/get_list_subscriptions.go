package subscriptions_service

import (
	"context"
	"fmt"

	"github.com/Vadick-do/Effective-Mobile/internal/core/domain"
)

func (s *SubscriptionsService) GetSubscriptions(
	ctx context.Context,
) ([]domain.Subscription, error) {
	subscriptions, err := s.subscriptionsRepository.GetSubscriptions(ctx)
	if err != nil {
		return nil, fmt.Errorf("get list subscriptions: %w", err)
	}

	return subscriptions, nil
}
