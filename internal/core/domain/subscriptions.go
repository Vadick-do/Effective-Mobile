package domain

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          uuid.UUID
	Version     int
	ServiceName string
	Price       int
	UserID      uuid.UUID
	StartDate   time.Time
	EndDate     *time.Time
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func NewSubscription(
	id uuid.UUID,
	version int,
	serviceName string,
	price int,
	userID uuid.UUID,
	startDate time.Time,
	endDate *time.Time,
	createdAt time.Time,
	updatedAt *time.Time,
) Subscription {
	return Subscription{
		ID:          id,
		Version:     version,
		ServiceName: serviceName,
		Price:       price,
		UserID:      userID,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
