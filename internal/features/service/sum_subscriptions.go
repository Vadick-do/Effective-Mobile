package subscriptions_service

import (
	"context"
	"fmt"

	"github.com/Vadick-do/Effective-Mobile/internal/core/domain"
	core_errors "github.com/Vadick-do/Effective-Mobile/internal/core/errors"
	"github.com/google/uuid"
)

func (s *SubscriptionsService) SumSubscriptions(
	ctx context.Context,
	from string,
	to string,
	userID *uuid.UUID,
	serviceName *string,
) (domain.Total, error) {
	if from > to {
		return domain.Total{}, fmt.Errorf("from must be <= to: %w", core_errors.ErrInvalidArgument)
	}

	total, err := s.subscriptionsRepository.SumSubscriptions(ctx, from, to, userID, serviceName)
	if err != nil {
		return domain.Total{}, fmt.Errorf("calculate total price: %w", err)
	}

	return total, nil
}
