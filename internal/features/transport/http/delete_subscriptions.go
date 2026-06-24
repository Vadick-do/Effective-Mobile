package subscriptions_transport_http

import (
	"net/http"

	core_logger "github.com/Vadick-do/Effective-Mobile/internal/core/logger"
	core_http_request "github.com/Vadick-do/Effective-Mobile/internal/core/transport/http/request"
	core_http_response "github.com/Vadick-do/Effective-Mobile/internal/core/transport/http/response"
)

func (h *SubscriptionsHTTPHandler) DeleteSubscription(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponse(log, rw)

	subID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get subscription id path value")
		return
	}

	if err := h.subscriptionsService.DeleteSubcsription(ctx, subID); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete subscription")
		return
	}

	responseHandler.NoContentResponse()
}
