package subscriptions_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *SubscriptionsService) DeleteSubcsription(
	ctx context.Context,
	subID uuid.UUID,
) error {
	if err := s.subscriptionsRepository.DeleteSubcsription(ctx, subID); err != nil {
		return fmt.Errorf("delete subscription: %w", err)
	}

	return nil
}
