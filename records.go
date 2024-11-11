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
	res, err := RequestPostRecord(m.CollectionID, new)

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

	res, err := RequestListRecords(m.CollectionID, pagination)

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
	res, err := RequestGetRecord(m.CollectionID, id)

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
	res, err := RequestUpdateRecord(m.CollectionID, id, target)

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
	res, err := RequestDeleteRecord(m.CollectionID, id)

	if err != nil {
		return err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return err
	}

	return nil
}
