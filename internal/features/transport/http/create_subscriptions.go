package subscriptions_transport_http

import (
	"encoding/json"
	"net/http"

	"github.com/Vadick-do/Effective-Mobile/internal/core/domain"
	core_logger "github.com/Vadick-do/Effective-Mobile/internal/core/logger"
	core_http_response "github.com/Vadick-do/Effective-Mobile/internal/core/transport/http/response"
)

type CreateSubscriptionResponse SubscriptionDTOResponse

// CreateSubscription   godoc
// @Summary             Создать подписку
// @Description         Создать новую подписку в системе.
// @Tags                subscriptions
// @Accept              json
// @Produce             json
// @Param               request body SubscriptionRequest true "CreateSubscription тело запроса"
// @Success             201 {object} CreateSubscriptionResponse "Успешно созданная подписка"
// @Failure             400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure             500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router              /subscriptions [post]
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
