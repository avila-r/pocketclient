package pocketclient_test

import (
	"testing"

	"github.com/avila-r/pocketclient"
)

func Test_Ping(t *testing.T) {
	client := pocketclient.New(pocketclient.Config{
		URL:      "http://127.0.0.1:8090",
		Email:    "avila.dev@outlook.com",
		Password: "1234567890",
	})

	if ok := client.Ping(); !ok {
		t.Errorf("failed - client health isn't ok")
	}
}
