package subscriptions_transport_http

import (
	"net/http"

	core_logger "github.com/Vadick-do/Effective-Mobile/internal/core/logger"
	core_http_request "github.com/Vadick-do/Effective-Mobile/internal/core/transport/http/request"
	core_http_response "github.com/Vadick-do/Effective-Mobile/internal/core/transport/http/response"
)

type GetSubscriptionResponse SubscriptionDTOResponse

func (h *SubscriptionsHTTPHandler) GetSubscription(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponse(log, rw)

	subID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get subscription id path value")
		return
	}

	subscriptionDomain, err := h.subscriptionsService.GetSubscription(ctx, subID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get subscription")
		return
	}

	response := GetSubscriptionResponse(subscriptionDTOFromDomain(subscriptionDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}
