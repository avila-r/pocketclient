package pocketclient_test

import (
	"testing"

	"github.com/avila-r/pocketclient"
	"github.com/avila-r/pocketclient/collections"
)

var (
	create_collection = func(t *testing.T) *collections.CollectionResponse {
		schema := collections.Schema{
			{
				Name: "name",
				Type: collections.Text,
			},
		}

		req, err := collections.New("testUsers").
			Schema(schema).
			Build()

		if err != nil {
			t.Errorf("failed to create collection request - %v", err.Error())
		}

		coll, err := pocketclient.Collections.Create(req)

		if err != nil {
			t.Logf("failed to create collection - %v", err.Error())
		}

		t.Cleanup(func() {
			if err := pocketclient.Collections.DeleteByID(coll.ID); err != nil {
				t.Errorf("failed to delete collection - %v", err.Error())
			}
		})

		return coll
	}
)

type User struct {
	Name string `json:"name"`
}

func Test_Insert(t *testing.T) {
	var (
		coll = create_collection(t)

		data = User{
			Name: "test",
		}
	)

	if err := pocketclient.Insert(coll.Name, data); err != nil {
		t.Errorf("failed to insert record - %v", err.Error())
	}
}
