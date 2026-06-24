package subscriptions_transport_http

import (
	"time"

	"github.com/Vadick-do/Effective-Mobile/internal/core/domain"
	core_http_response "github.com/Vadick-do/Effective-Mobile/internal/core/transport/http/response"
	"github.com/google/uuid"
)

type SubscriptionDTOResponse struct {
	ID          uuid.UUID  `json:"id"            example:"a3f2e8d1-9b4c-4c7f-b2d6-1e5a8c9f0b3d"`
	ServiceName string     `json:"service_name"  example:"Yandex Plus"`
	Price       int        `json:"price"         example:"400"`
	UserID      uuid.UUID  `json:"user_id"       example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	StartDate   string     `json:"start_date"    example:"07-2025"`
	EndDate     *string    `json:"end_date"      example:"08-2025"`
	CreatedAt   time.Time  `json:"created_at"    example:"2026-02-26T10:30:00Z"`
	UpdatedAt   *time.Time `json:"updated_at"    example:"2026-06-24T16:15:18.62982Z"`
}

type SubscriptionsTotalPrice struct {
	Total int `json:"total" example:"1000"`
}

type SubscriptionRequest struct {
	ServiceName string    `json:"service_name"       example:"Yandex Plus"`
	Price       int       `json:"price"              example:"400"`
	UserID      uuid.UUID `json:"user_id"            example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	StartDate   string    `json:"start_date"         example:"07-2025"`
	EndDate     *string   `json:"end_date,omitempty" example:"08-2025"`
}

func subscriptionDTOFromDomain(sub domain.Subscription) SubscriptionDTOResponse {
	startDate := core_http_response.FormatDate(sub.StartDate)
	sub.StartDate = startDate
	if sub.EndDate != nil {
		endDate := core_http_response.FormatDate(*sub.EndDate)
		sub.EndDate = &endDate
	}
	return SubscriptionDTOResponse{
		ID:          sub.ID,
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserID:      sub.UserID,
		StartDate:   sub.StartDate,
		EndDate:     sub.EndDate,
		CreatedAt:   sub.CreatedAt,
		UpdatedAt:   sub.UpdatedAt,
	}
}

func subscriptionsDTOFromDomains(subs []domain.Subscription) []SubscriptionDTOResponse {
	subDTO := make([]SubscriptionDTOResponse, len(subs))

	for i, sub := range subs {
		subDTO[i] = subscriptionDTOFromDomain(sub)
	}

	return subDTO
}

func subscriptionsTotalPrice(total domain.Total) SubscriptionsTotalPrice {
	return SubscriptionsTotalPrice{
		Total: total.Total,
	}
}
