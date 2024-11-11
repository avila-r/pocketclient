package pocketclient_test

import (
	"testing"

	"github.com/avila-r/pocketclient"
	"github.com/avila-r/pocketclient/collections"
)

var (
	collection_request = collections.CollectionRequest{
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

	test_create_collection = func(t *testing.T, req ...collections.CollectionRequest) *collections.CollectionResponse {
		r := collection_request
		if len(req) > 0 {
			r = req[0]
		}

		id := r.Name

		t.Cleanup(func() {
			test_delete_collection(id)
		})

		if collection_exists(id) {
			return test_find_collection(t, id)
		}

		res, err := pocketclient.Collections.Create(&r)

		if err != nil {
			t.Errorf("failed to create collection - %v", err.Error())
		}

		return res
	}

	test_list_collection = func(t *testing.T) *pocketclient.Pagination[collections.CollectionResponse] {
		res, err := pocketclient.Collections.ListAll()

		if err != nil {
			t.Errorf("failed to list collections - %v", err.Error())
		}

		return res
	}

	test_find_collection = func(t *testing.T, id string) *collections.CollectionResponse {
		res, err := pocketclient.Collections.GetByID(id)

		if err != nil {
			t.Errorf("failed to get collection by id %v - %v", id, err.Error())
		}

		return res
	}

	test_delete_collection = func(id string) error {
		return pocketclient.Collections.DeleteByID(id)
	}

	collection_exists = func(id string) bool {
		_, err := pocketclient.Collections.GetByID(id)

		return err == nil
	}
)

func Test_Create(t *testing.T) {
	_ = test_create_collection(t)
}

func Test_Delete(t *testing.T) {
	res := test_create_collection(t)

	if err := test_delete_collection(res.ID); err != nil {
		t.Errorf("failed to delete collection - %v", err.Error())
	}
}

func Test_List(t *testing.T) {
	_ = test_create_collection(t)

	test_list_collection(t)
}

func Test_GetByID(t *testing.T) {
	test_find_collection(t, test_create_collection(t).ID)
}

func Test_Builder(t *testing.T) {
	r, err := collections.New(collection_request.Name).
		Schema(collections.Schema{
			{
				Name: "name",
				Type: collections.Text,
			},
			{
				Name: "surname",
				Type: collections.Text,
			},
		}).
		Build()

	if err != nil {
		t.Errorf("failed to build request - %v", err.Error())
	}

	_ = test_create_collection(t, *r)
}
