package pocketclient_test

import (
	"testing"

	"github.com/avila-r/pocketclient"
	"github.com/avila-r/pocketclient/internal/tests"
)

func Test_ListAll(t *testing.T) {
	client := tests.PocketClient

	// or pocketclient.Admin.ListAll()
	pagination, err := client.Admin().ListAll()

	if err != nil {
		t.Errorf("unable to list admins - %v", err.Error())
	}

	t.Log(pagination)
}

func Test_RequestPasswordReset(t *testing.T) {
	client := tests.PocketClient

	current := client.PocketBase.Credentials.Email
	if err := pocketclient.Admin.RequestPasswordReset(current); err != nil {
		t.Errorf("unable to request password reset - %v", err.Error())
	}
}
