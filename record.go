package pocketclient

import "github.com/avila-r/pocketclient/validation"

func Insert(collection string, data any) error {
	if !Client.IsAuthenticated() {
		return ErrNotAuthenticated
	}

	res, err := Client.Resty.R().
		SetHeader(HeaderAuthorizationToken()).
		SetBody(data).
		Post(Client.PocketBase.URL + EndpointCollection(collection))

	if err != nil {
		return err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return err
	}

	if err := Unmarshal(res.Body(), &data); err != nil {
		return err
	}

	return nil
}

func List[T any](collection string, p ...PaginationParams) (*Pagination[T], error) {
	return Collection[T](collection).
		List(p...)
}

type Query struct {
	Collection string
	ID         string
}

func Find[T any](q Query) (*T, error) {
	return Collection[T](q.Collection).
		Find(q.ID)
}

func FindIn[T any](collection string, id string) (*T, error) {
	return Collection[T](collection).
		Find(id)
}

func Fetch(q Query, to any) error {
	if !Client.IsAuthenticated() {
		return ErrNotAuthenticated
	}

	res, err := Client.Resty.R().
		SetHeader(HeaderAuthorizationToken()).
		Get(Client.Resty.BaseURL + EndpointCollection(q.Collection) + "/" + q.ID)

	if err != nil {
		return err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return err
	}

	if err := Unmarshal(res.Body(), &to); err != nil {
		return err
	}

	return nil
}

func FetchIn(collection, id string, to any) error {
	query := Query{
		Collection: collection,
		ID:         id,
	}

	return Fetch(query, &to)
}
