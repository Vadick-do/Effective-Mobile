package subscriptions_transport_http

import (
	"encoding/json"
	"net/http"

	"github.com/Vadick-do/Effective-Mobile/internal/core/domain"
	core_logger "github.com/Vadick-do/Effective-Mobile/internal/core/logger"
	core_http_request "github.com/Vadick-do/Effective-Mobile/internal/core/transport/http/request"
	core_http_response "github.com/Vadick-do/Effective-Mobile/internal/core/transport/http/response"
)

type PutSubcsriptionResponse SubscriptionDTOResponse

func (h *SubscriptionsHTTPHandler) PutSubsciption(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponse(log, rw)

	subID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get subscription id path value")
		return
	}

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

	subDomain, err := h.subscriptionsService.PutSubscription(ctx, subID, subscriptionDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to put subscription")
		return
	}

	response := PutSubcsriptionResponse(subscriptionDTOFromDomain(subDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}
