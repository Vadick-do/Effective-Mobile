package subscriptions_transport_http

import (
	"context"
	"net/http"

	"github.com/Vadick-do/Effective-Mobile/internal/core/domain"
	core_http_server "github.com/Vadick-do/Effective-Mobile/internal/core/transport/http/server"
	"github.com/google/uuid"
)

type SubscriptionsHTTPHandler struct {
	subscriptionsService SubscriptionsService
}

type SubscriptionsService interface {
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

func NewSubscriptionsHTTPHandler(
	subscriptionsService SubscriptionsService,
) *SubscriptionsHTTPHandler {
	return &SubscriptionsHTTPHandler{
		subscriptionsService: subscriptionsService,
	}
}

func (h *SubscriptionsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/subscriptions",
			Handler: h.CreateSubscription,
		},
		{
			Method:  http.MethodGet,
			Path:    "/subscriptions/{id}",
			Handler: h.GetSubscription,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/subscriptions/{id}",
			Handler: h.DeleteSubscription,
		},
		{
			Method:  http.MethodGet,
			Path:    "/subscriptions",
			Handler: h.GetSubscriptions,
		},
		{
			Method:  http.MethodPut,
			Path:    "/subscriptions/{id}",
			Handler: h.PutSubsciption,
		},
		{
			Method:  http.MethodGet,
			Path:    "/subscriptions/total",
			Handler: h.Sumsubscriptions,
		},
	}
}
