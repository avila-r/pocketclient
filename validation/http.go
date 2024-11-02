package validation

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

var (
	ErrBadRequest    = errors.New("something went wrong while processing your request")
	ErrNotAuthorized = errors.New("the request requires valid authorization token")
	ErrNotAllowed    = errors.New("not allowed to perform request")
	ErrNotFound      = errors.New("requested resource not found")
)

func VerifyResponse(r *resty.Response) error {
	switch r.StatusCode() {
	case 400:
		return ErrBadRequest
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
