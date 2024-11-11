package pocketclient

import (
	"github.com/go-resty/resty/v2"

	"github.com/avila-r/pocketclient/collections"
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
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			SetQueryParams(pagination.ToQueryParams()).
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
	RequestPostRecord = func(collection string, new any, client ...*PocketClient) (*resty.Response, error) {
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
			Post(Client.PocketBase.URL + EndpointCollection(collection))
	}

	RequestListRecords = func(collection string, pagination PaginationParams, client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		if !c.IsAuthenticated() {
			return nil, ErrNotAuthenticated
		}

		return c.Resty.R().
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			SetQueryParams(pagination.ToQueryParams()).
			Get(Client.PocketBase.URL + EndpointCollection(collection))
	}

	RequestGetRecord = func(collection string, id string, client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		if !c.IsAuthenticated() {
			return nil, ErrNotAuthenticated
		}

		return c.Resty.R().
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			Get(Client.Resty.BaseURL + EndpointCollection(collection) + "/" + id)
	}

	RequestUpdateRecord = func(collection string, id string, patch any, client ...*PocketClient) (*resty.Response, error) {
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
			Patch(Client.PocketBase.URL + EndpointCollection(collection) + "/" + id)
	}

	RequestDeleteRecord = func(collection string, id string, client ...*PocketClient) (*resty.Response, error) {
		c := Client
		if len(client) > 0 {
			c = client[0]
		}

		if !c.IsAuthenticated() {
			return nil, ErrNotAuthenticated
		}

		return c.Resty.R().
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			Delete(Client.PocketBase.URL + EndpointCollection(collection) + "/" + id)
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
			SetHeader(HeaderAuthorizationTokenFrom(c)).
			SetQueryParams(pagination.ToQueryParams()).
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
