package pocketclient

import (
	"github.com/avila-r/pocketclient/collections"

	"github.com/go-resty/resty/v2"
)

var (
	RequestPostCollection = func(request *collections.CollectionRequest, client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		return c.Resty.R().
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			SetBody(request).
			Post(Client.PocketBase.URL + EndpointCollections)
	}

	RequestListCollections = func(pagination PaginationParams, client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		return c.Resty.R().
			SetQueryParams(pagination.ToQueryParams()).
			Get(Client.PocketBase.URL + EndpointCollections)
	}

	RequestGetCollection = func(id string, client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		return c.Resty.R().
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			Get(Client.Resty.BaseURL + EndpointCollections + "/" + id)
	}

	RequestUpdateCollection = func(id string, new *collections.CollectionPatch, client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		return c.Resty.R().
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			SetBody(new).
			Patch(Client.PocketBase.URL + EndpointCollections + "/" + id)
	}

	RequestDeleteCollection = func(id string, client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		return c.Resty.R().
			SetHeader(HeaderAuthorizationToken()).
			Delete(Client.PocketBase.URL + EndpointCollections + "/" + id)
	}
)
