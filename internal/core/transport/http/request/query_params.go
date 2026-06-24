package core_http_request

import (
	"fmt"
	"net/http"

	core_errors "github.com/Vadick-do/Effective-Mobile/internal/core/errors"
	core_http_response "github.com/Vadick-do/Effective-Mobile/internal/core/transport/http/response"
	"github.com/google/uuid"
)

func GetUUIDQueryParam(r *http.Request, key string) (*uuid.UUID, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	val, err := uuid.Parse(param)
	if err != nil {
		return nil, fmt.Errorf(
			"param='%s' by key='%s' not a valid uuid: %v: %w",
			param,
			key,
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	return &val, nil
}

func GetDateQueryParam(r *http.Request, key string) (string, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return "", fmt.Errorf("missing required query params to/from '%s': %w", key, core_errors.ErrInvalidArgument)
	}

	date := core_http_response.FormatDate(param)

	return date, nil
}

func GetServiceNameQueryParam(r *http.Request, key string) *string {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil
	}

	return &param
}
