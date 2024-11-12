package validation

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

var (
	ErrBadRequest = Failure{
		Code:    400,
		Message: "something went wrong while processing your request",
	}

	ErrNotAuthorized = Failure{
		Code:    401,
		Message: "the request requires valid authorization token",
	}

	ErrNotAllowed = Failure{
		Code:    403,
		Message: "not allowed to perform request",
	}

	ErrNotFound = Failure{
		Code:    404,
		Message: "requested resource not found",
	}
)

func VerifyResponse(r *resty.Response) error {
	switch r.StatusCode() {
	case 400:
		return CatchError(r)
	case 401:
		return ErrNotAuthorized
	case 403:
		return ErrNotAllowed
	case 404:
		return ErrNotFound
	default:
		return nil
	}
}

func CatchError(r *resty.Response) (f error) {
	if err := json.Unmarshal(r.Body(), &f); err != nil {
		return ErrBadRequest
	}

	return
}
