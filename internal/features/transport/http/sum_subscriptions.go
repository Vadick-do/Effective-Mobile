package subscriptions_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Vadick-do/Effective-Mobile/internal/core/logger"
	core_http_request "github.com/Vadick-do/Effective-Mobile/internal/core/transport/http/request"
	core_http_response "github.com/Vadick-do/Effective-Mobile/internal/core/transport/http/response"
	"github.com/google/uuid"
)

type TotalPriceResponse SubscriptionsTotalPrice

// SumSubscriptions     godoc
// @Summary             Получение общей суммы подписок
// @Description         Получение общей суммы подписок за указанный период с опциональной фильтрацией по id пользователя и названию сервиса
// @Tags                subscriptions
// @Produce             json
// @Param               from          query string true  "Дата начала периода в формате MM-YYYY"
// @Param               to            query string true  "Дата окончания периода в формате MM-YYYY"
// @Param               user_id       query string false "Фильтрация подписок по ID пользователя" Format(uuid)
// @Param               service_name  query string false "Фильтрация подписок по названию сервиса"
// @Success             200 {object}  TotalPriceResponse "Общая стоимость подписок"
// @Failure             400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure             500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router              /subscriptions/total [get]
func (h *SubscriptionsHTTPHandler) Sumsubscriptions(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponse(log, rw)

	from, to, userID, serviceName, err := getUserIDServiceNameFromToQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get fom/to/user_id/service_name query params")
		return
	}

	totalDomain, err := h.subscriptionsService.SumSubscriptions(ctx, from, to, userID, serviceName)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to calculate the total amount of subscriptions")
		return
	}

	response := TotalPriceResponse(subscriptionsTotalPrice(totalDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}

func getUserIDServiceNameFromToQueryParams(r *http.Request) (string, string, *uuid.UUID, *string, error) {
	const (
		userIDQueryParamKey      = "user_id"
		fromQueryParamKey        = "from"
		toQueryParamKey          = "to"
		serviceNamequeryParamKey = "service_name"
	)

	userID, err := core_http_request.GetUUIDQueryParam(r, userIDQueryParamKey)
	if err != nil {
		return "", "", nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}

	from, err := core_http_request.GetDateQueryParam(r, fromQueryParamKey)
	if err != nil {
		return "", "", nil, nil, fmt.Errorf("get 'from' query param: %w", err)
	}

	to, err := core_http_request.GetDateQueryParam(r, toQueryParamKey)
	if err != nil {
		return "", "", nil, nil, fmt.Errorf("get 'to' query param: %w", err)
	}

	serviceName := core_http_request.GetServiceNameQueryParam(r, serviceNamequeryParamKey)

	return from, to, userID, serviceName, nil
}
