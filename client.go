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
	Client *PocketClient
)

func New(config Config) *PocketClient {
	c := &PocketClient{
		Resty: resty.New(),
		PocketBase: &PocketBase{
			URL: "http://127.0.0.1:8090",
		},
	}

	if config.URL != "" {
		c.PocketBase.URL = config.URL
	}

	Client = c

	auth, err := Admin.auth(config.Email, config.Password)

	if err != nil {
		log.Fatalf("failed to authenticate first admin - %v", err.Error())
	}

	c.PocketBase.Admins = append(c.PocketBase.Admins, auth)

	return Client
}
