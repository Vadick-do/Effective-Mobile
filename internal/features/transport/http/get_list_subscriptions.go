package subscriptions_transport_http

import (
	"net/http"

	core_logger "github.com/Vadick-do/Effective-Mobile/internal/core/logger"
	core_http_response "github.com/Vadick-do/Effective-Mobile/internal/core/transport/http/response"
)

type GetSubscriptionsResponse []SubscriptionDTOResponse

// GetSubscriptions     godoc
// @Summary             Список подписок
// @Description         Просмотр списка всех подписок
// @Tags                subscriptions
// @Produce             json
// @Success             200 {object} GetSubscriptionsResponse "Список подписок"
// @Failure             500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router              /subscriptions [get]
func (h *SubscriptionsHTTPHandler) GetSubscriptions(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponse(log, rw)

	subsDomains, err := h.subscriptionsService.GetSubscriptions(ctx)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get list subscriptions")
		return
	}

	response := GetSubscriptionsResponse(subscriptionsDTOFromDomains(subsDomains))
	responseHandler.JSONResponse(response, http.StatusOK)
}
