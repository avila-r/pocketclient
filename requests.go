package pocketclient

import (
	"github.com/avila-r/pocketclient/collections"

	"github.com/go-resty/resty/v2"
)

var (
	RequestPostAdmin = func(request AdminRequest, client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		if !c.IsAuthenticated() {
			return nil, ErrNotAuthenticated
		}

		return c.Resty.R().
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			SetBody(request).
			Post(Client.PocketBase.URL + EndpointAdmins)
	}

	RequestListAdmins = func(pagination PaginationParams, client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		if !c.IsAuthenticated() {
			return nil, ErrNotAuthenticated
		}

		return c.Resty.R().
			SetQueryParams(pagination.ToQueryParams()).
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			Get(Client.PocketBase.URL + EndpointAdmins)
	}

	RequestGetAdmin = func(id string, client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		if !c.IsAuthenticated() {
			return nil, ErrNotAuthenticated
		}

		return c.Resty.R().
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			Get(Client.Resty.BaseURL + EndpointAdmins + "/" + id)
	}

	RequestUpdateAdmin = func(id string, patch AdminRequest, client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		if !c.IsAuthenticated() {
			return nil, ErrNotAuthenticated
		}

		return c.Resty.R().
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			SetBody(patch).
			Patch(Client.PocketBase.URL + EndpointAdmins + "/" + id)
	}

	RequestDeleteAdmin = func(id string, client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		if !c.IsAuthenticated() {
			return nil, ErrNotAuthenticated
		}

		return c.Resty.R().
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			Delete(Client.PocketBase.URL + EndpointAdmins + "/" + id)
	}
)

var (
	RequestPostCollection = func(request *collections.CollectionRequest, client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		if !c.IsAuthenticated() {
			return nil, ErrNotAuthenticated
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
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			SetQueryParams(pagination.ToQueryParams()).
			Get(Client.PocketBase.URL + EndpointCollections)
	}

	RequestGetCollection = func(id string, client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		if !c.IsAuthenticated() {
			return nil, ErrNotAuthenticated
		}

		return c.Resty.R().
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			Get(Client.PocketBase.URL + EndpointCollections + "/" + id)
	}

	RequestUpdateCollection = func(id string, new *collections.CollectionPatch, client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		if !c.IsAuthenticated() {
			return nil, ErrNotAuthenticated
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

		if !c.IsAuthenticated() {
			return nil, ErrNotAuthenticated
		}

		return c.Resty.R().
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			Delete(Client.PocketBase.URL + EndpointCollections + "/" + id)
	}
)

var (
	RequestListLogs = func(pagination PaginationParams, client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		if !c.IsAuthenticated() {
			return nil, ErrNotAuthenticated
		}

		return c.Resty.R().
			SetQueryParams(pagination.ToQueryParams()).
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			Get(c.PocketBase.URL + EndpointLogs)
	}

	RequestGetLog = func(id string, client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		if !c.IsAuthenticated() {
			return nil, ErrNotAuthenticated
		}

		return c.Resty.R().
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			Get(c.PocketBase.URL + EndpointLogs + "/" + id)
	}

	RequestGetLoggingStats = func(client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		if !c.IsAuthenticated() {
			return nil, ErrNotAuthenticated
		}

		return c.Resty.R().
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			Get(c.PocketBase.URL + EndpointLogsStats)
	}
)
