package subscriptions_transport_http

import (
	"encoding/json"
	"net/http"

	"github.com/Vadick-do/Effective-Mobile/internal/core/domain"
	core_logger "github.com/Vadick-do/Effective-Mobile/internal/core/logger"
	core_http_response "github.com/Vadick-do/Effective-Mobile/internal/core/transport/http/response"
)

type CreateSubscriptionResponse SubscriptionDTOResponse

func (h *SubscriptionsHTTPHandler) CreateSubscription(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponse(log, rw)

	var request SubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode HTTP request")
		return
	}

	startDate := core_http_response.FormatDate(request.StartDate)
	request.StartDate = startDate

	if request.EndDate != nil {
		endDate := core_http_response.FormatDate(*request.EndDate)
		request.EndDate = &endDate
	}

	subscriptionDomain := domain.NewSubscriptionUninitialized(
		request.ServiceName,
		request.Price,
		request.UserID,
		request.StartDate,
		request.EndDate,
	)

	subscriptionDomain, err := h.subscriptionsService.CreateSubscription(ctx, subscriptionDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create subscription")
		return
	}

	response := CreateSubscriptionResponse(subscriptionDTOFromDomain(subscriptionDomain))
	responseHandler.JSONResponse(response, http.StatusCreated)
}
