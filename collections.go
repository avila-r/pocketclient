package pocketclient

import (
	"github.com/avila-r/pocketclient/collections"
	"github.com/avila-r/pocketclient/validation"
)

func (m *ModuleCollections) New(new *collections.CollectionRequest) (*collections.CollectionResponse, error) {
	res, err := RequestPostCollection(new)

	if err != nil {
		return nil, err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return nil, err
	}

	created := collections.CollectionResponse{}
	if err := Unmarshal(res.Body(), &created); err != nil {
		return nil, err
	}

	return &created, nil
}

func (m *ModuleCollections) ListAll(p ...PaginationParams) (*Pagination[collections.CollectionResponse], error) {
	pagination := PaginationParams{}
	if len(p) > 0 {
		pagination = p[0]
	}

	res, err := RequestListCollections(pagination)

	if err != nil {
		return nil, err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return nil, err
	}

	collections := Pagination[collections.CollectionResponse]{}
	if err := Unmarshal(res.Body(), &collections); err != nil {
		return nil, err
	}

	return &collections, nil
}

func (m *ModuleCollections) GetByID(id string) (*collections.CollectionResponse, error) {
	res, err := RequestGetCollection(id)

	if err != nil {
		return nil, err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return nil, err
	}

	collection := collections.CollectionResponse{}
	if err := Unmarshal(res.Body(), &collection); err != nil {
		return nil, err
	}

	return &collection, nil
}

func (m *ModuleCollections) Update(id string, new *collections.CollectionPatch) (*collections.CollectionResponse, error) {
	res, err := RequestUpdateCollection(id, new)

	if err != nil {
		return nil, err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return nil, err
	}

	updated := collections.CollectionResponse{}
	if err := Unmarshal(res.Body(), &updated); err != nil {
		return nil, err
	}

	return &updated, nil
}

func (m *ModuleCollections) DeleteByID(id string) error {
	res, err := RequestDeleteCollection(id)

	if err != nil {
		return err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return err
	}

	return nil
}
