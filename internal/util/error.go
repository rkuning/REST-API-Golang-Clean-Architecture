package util

import (
	"fmt"
	"net/http"
	"rest-api-golang-clean-code/internal/model"
)

type ApiError model.WebResponse[error]

func (e ApiError) Error() string {
	return fmt.Sprintf("api error : %d", e.Code)
}

func NewApiError(code int, message string) ApiError {
	return ApiError{Code: code, Message: message}
}

func ValidateError(err error) ApiError {
	return NewApiError(http.StatusUnprocessableEntity, ExactError(err))
}

func ConflictError(message string) ApiError {
	if message == "" {
		message = "data already exist"
	}
	return NewApiError(http.StatusConflict, message)
}

func UnauthorizedError() ApiError {
	return NewApiError(http.StatusUnauthorized, "unauthorized access")
}

func NotFoundError(message string) ApiError {
	if message == "" {
		message = "data not found"
	}
	return NewApiError(http.StatusNotFound, message)
}

func BadRequestError(message string) ApiError {
	if message == "" {
		message = "bad request"
	}
	return NewApiError(http.StatusBadRequest, message)
}

func UnprocessableError(message string) ApiError {
	if message == "" {
		message = "unprocessable entity"
	}
	return NewApiError(http.StatusUnprocessableEntity, message)
}

func InternalServerError() ApiError {
	return NewApiError(http.StatusInternalServerError, "internal server error")
}

func TimeoutError() ApiError {
	return NewApiError(http.StatusGatewayTimeout, "request timeout")
}
