package subscriptions_transport_http

import (
	"encoding/json"
	"net/http"

	"github.com/Vadick-do/Effective-Mobile/internal/core/domain"
	core_logger "github.com/Vadick-do/Effective-Mobile/internal/core/logger"
	core_http_request "github.com/Vadick-do/Effective-Mobile/internal/core/transport/http/request"
	core_http_response "github.com/Vadick-do/Effective-Mobile/internal/core/transport/http/response"
)

type PutSubscriptionResponse SubscriptionDTOResponse

// PutSubscription     godoc
// @Summary            Обновить запись о подписке полностью
// @Description        Обновляет информацию об уже существующей в системе подписке.
// @Description        Тело запроса должно содержать все поля подписки (как при создании).
// @Description        Поле `end_date` опционально; если оно не передано или равно null, дата окончания сбрасывается.
// @Tags               subscriptions
// @Accept             json
// @Produce            json
// @Param              id path string true "ID обновляемой подписки" Format(uuid)
// @Param              request body SubscriptionRequest true "PutSubscription тело запроса"
// @Success            200 {object} PutSubscriptionResponse "Успешно измененная подписка"
// @Failure            400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure            404 {object} core_http_response.ErrorResponse "Subscription not found"
// @Failure            500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router             /subscriptions/{id} [put]
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

	response := PutSubscriptionResponse(subscriptionDTOFromDomain(subDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}
