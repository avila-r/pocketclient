package pocketclient

import (
	"log"

	"github.com/go-resty/resty/v2"
)

type (
	PocketClient struct {
		Resty      *resty.Client
		PocketBase *PocketBase
	}

	Config struct {
		URL string

		Email    string
		Password string
	}
)

var (
	Client = &PocketClient{
		Resty: resty.New(),
		PocketBase: &PocketBase{
			URL: "http://127.0.0.1:8090",
		},
	}
)

func New(config Config) *PocketClient {
	if config.URL != "" {
		Client.PocketBase.URL = config.URL
	}

	auth, err := Admin.Auth(config.Email, config.Password)

	if err != nil {
		log.Printf("[pocketclient] failed to authenticate (was first admin created?) - %v", err.Error())
	}

	Client.PocketBase.Credentials = auth

	return Client
}
