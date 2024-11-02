package pocketclient

import (
	"github.com/avila-r/pocketclient/validation"
)

type RecordModule[T any] struct {
	CollectionID string
}

func Collection[T any](id string) *RecordModule[T] {
	return &RecordModule[T]{
		CollectionID: id,
	}
}

func (m *RecordModule[T]) Insert(new T) (*T, error) {
	if !Client.IsAuthenticated() {
		return nil, ErrNotAuthenticated
	}

	res, err := Client.Resty.R().
		SetHeader(HeaderAuthorizationToken()).
		SetBody(new).
		Post(Client.PocketBase.URL + EndpointCollection(m.CollectionID))

	if err != nil {
		return nil, err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return nil, err
	}

	var created T
	if err := Unmarshal(res.Body(), &created); err != nil {
		return nil, err
	}

	return &created, err
}

func (m *RecordModule[T]) List(p ...PaginationParams) (*Pagination[T], error) {
	pagination := PaginationParams{}
	if len(p) > 0 {
		pagination = p[0]
	}

	if !Client.IsAuthenticated() {
		return nil, ErrNotAuthenticated
	}

	res, err := Client.Resty.R().
		SetQueryParams(pagination.ToQueryParams()).
		Get(Client.PocketBase.URL + EndpointCollection(m.CollectionID))

	if err != nil {
		return nil, err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return nil, err
	}

	records := Pagination[T]{}
	if err := Unmarshal(res.Body(), &records); err != nil {
		return nil, err
	}

	return &records, nil
}

func (m *RecordModule[T]) Find(id string) (*T, error) {
	if !Client.IsAuthenticated() {
		return nil, ErrNotAuthenticated
	}

	res, err := Client.Resty.R().
		SetHeader(HeaderAuthorizationToken()).
		Get(Client.Resty.BaseURL + EndpointCollection(m.CollectionID) + "/" + id)

	if err != nil {
		return nil, err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return nil, err
	}

	var record T
	if err := Unmarshal(res.Body(), &record); err != nil {
		return nil, err
	}

	return &record, nil
}

func (m *RecordModule[T]) Update(id string, target *T) error {
	if !Client.IsAuthenticated() {
		return ErrNotAuthenticated
	}

	res, err := Client.Resty.R().
		SetHeader(HeaderAuthorizationToken()).
		SetBody(*target).
		Patch(Client.PocketBase.URL + EndpointCollection(m.CollectionID) + "/" + id)

	if err != nil {
		return err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return err
	}

	if err := Unmarshal(res.Body(), target); err != nil {
		return err
	}

	return nil
}

func (m *RecordModule[T]) Delete(id string) error {
	if !Client.IsAuthenticated() {
		return ErrNotAuthenticated
	}

	res, err := Client.Resty.R().
		SetHeader(HeaderAuthorizationToken()).
		Delete(Client.PocketBase.URL + EndpointCollection(m.CollectionID) + "/" + id)

	if err != nil {
		return err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return err
	}

	return nil
}
