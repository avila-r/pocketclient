package pocketclient_test

import (
	"testing"

	"github.com/avila-r/pocketclient"
	"github.com/avila-r/pocketclient/collections"
)

var (
	request = collections.CollectionRequest{
		Name: "collectionBase",
		Type: collections.TypeBase,
		Schema: []collections.CollectionRequestField{
			{
				Name: "name",
				Type: collections.Text,
			},
			{
				Name: "surname",
				Type: collections.Text,
			},
		},
	}

	create = func(t *testing.T) *collections.CollectionResponse {
		id := request.Name

		t.Cleanup(func() {
			delete(id)
		})

		if exists(id) {
			return find(t, id)
		}

		res, err := pocketclient.Collections.Create(&request)

		if err != nil {
			t.Errorf("failed to create collection - %v", err.Error())
		}

		return res
	}

	list = func(t *testing.T) *pocketclient.Pagination[collections.CollectionResponse] {
		res, err := pocketclient.Collections.ListAll()

		if err != nil {
			t.Errorf("failed to list collections - %v", err.Error())
		}

		return res
	}

	find = func(t *testing.T, id string) *collections.CollectionResponse {
		res, err := pocketclient.Collections.GetByID(id)

		if err != nil {
			t.Errorf("failed to get collection by id %v - %v", id, err.Error())
		}

		return res
	}

	delete = func(id string) error {
		return pocketclient.Collections.DeleteByID(id)
	}

	exists = func(id string) bool {
		_, err := pocketclient.Collections.GetByID(id)

		return err == nil
	}
)

func Test_Create(t *testing.T) {
	_ = create(t)
}

func Test_Delete(t *testing.T) {
	res := create(t)

	if err := delete(res.ID); err != nil {
		t.Errorf("failed to delete collection - %v", err.Error())
	}
}

func Test_List(t *testing.T) {
	_ = create(t)

	list(t)
}

func Test_GetByID(t *testing.T) {
	find(t, create(t).ID)
}
