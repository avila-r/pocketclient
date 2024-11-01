package pocketclient

func (c *PocketClient) IsAuthenticated() bool {
	return c.PocketBase.Credentials.Token != ""
}

func (c *PocketClient) Ping() bool {
	res, err := c.Resty.R().
		Head(c.PocketBase.URL + EndpointHealthCheck)

	if err != nil || res.StatusCode() != 200 {
		return false
	}

	return true
}
