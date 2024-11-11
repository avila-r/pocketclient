package pocketclient

import (
	"regexp"

	"github.com/avila-r/pocketclient/collections"
	"github.com/avila-r/pocketclient/validation"
)

func (m *ModuleCollections) New(name string) *collections.CollectionBuilder {
	return collections.New(name)
}

func (m *ModuleCollections) Create(new *collections.CollectionRequest) (*collections.CollectionResponse, error) {
	if err := m.Validate(new); err != nil {
		return nil, err
	}

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

func (m *ModuleCollections) Validate(r *collections.CollectionRequest) error {
	if r.Name == "" {
		return Error("missing required value (name)")
	}

	if !isCamelCase(r.Name) {
		return Error("collection's name must be in camel case pattern")
	}

	if string(r.Type) == "" {
		return Error("missing required value (type)")
	}

	if r.Type == collections.TypeBase && r.Schema == nil {
		return Error("schema is required for base collections")
	}

	if r.Type == collections.TypeView && r.Options.ViewOptions == nil {
		return Error("query is required for view collections")
	}

	if r.Type == collections.TypeView && r.Options.ViewOptions.Query == "" {
		return Error("query is required for view collections")
	}

	return nil
}

func isCamelCase(s string) bool {
	r, _ := regexp.Compile(`^[a-z]+(?:[A-Z][a-z]*)*$`)

	return r.MatchString(s)
}
