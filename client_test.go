package pocketclient_test

import (
	"net/http"
	"testing"

	"github.com/avila-r/pocketclient"
)

func Test_FirstAdmin(t *testing.T) {
	client := pocketclient.New(pocketclient.Config{
		URL:      "http://127.0.0.1:8090",
		Email:    "avila.dev@outlook.com",
		Password: "1234567890",
	})

	t.Logf("token - %v", client.PocketBase.FirstAdmin().Token)

	response, err := client.Resty.R().
		SetHeader(pocketclient.HeaderAuthorizationToken()).
		Get(client.PocketBase.URL + pocketclient.EndpointAdmins)

	if err != nil {
		t.Errorf("failed - %v", err)
	}

	if expected, got := http.StatusOK, response.StatusCode(); got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
