package domain

import (
	"time"

	"github.com/google/uuid"
)

func generateID() uuid.UUID {
	return uuid.New()
}

type Total struct {
	Total int
}

type Subscription struct {
	ID          uuid.UUID
	ServiceName string
	Price       int
	UserID      uuid.UUID
	StartDate   string
	EndDate     *string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func NewSubscription(
	id uuid.UUID,
	serviceName string,
	price int,
	userID uuid.UUID,
	startDate string,
	endDate *string,
	createdAt time.Time,
	updatedAt *time.Time,
) Subscription {
	return Subscription{
		ID:          id,
		ServiceName: serviceName,
		Price:       price,
		UserID:      userID,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func NewSubscriptionUninitialized(
	serviceName string,
	price int,
	userID uuid.UUID,
	startDate string,
	endDate *string,
) Subscription {
	id := generateID()
	return Subscription{
		ID:          id,
		ServiceName: serviceName,
		Price:       price,
		UserID:      userID,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}
}
