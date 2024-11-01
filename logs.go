package pocketclient

type Log struct {
	ID string `json:"id"`

	CreatedAt string `json:"created"`
	UpdatedAt string `json:"updated"`

	Message string `json:"message"`
	Level   int    `json:"level"`

	Details struct {
		Auth      string  `json:"auth"`
		ExecTime  float64 `json:"execTime"`
		Method    string  `json:"method"`
		Referer   string  `json:"referer"`
		RemoteIp  string  `json:"remoteIp"`
		Status    int     `json:"status"`
		Type      string  `json:"type"`
		URL       string  `json:"url"`
		UserAgent string  `json:"userAgent"`
		UserIp    string  `json:"userIp"`
	} `json:"data"`
}

func (c *PocketClient) GetLogs(p ...PaginationParams) (*Pagination[Log], error) {
	params := PaginationParams{}
	if len(p) > 0 {
		params = p[0]
	}

	if !c.IsAuthenticated() {
		return nil, ErrNotAuthenticated
	}

	res, err := c.Resty.R().
		SetQueryParams(params.ToQueryParams()).
		SetHeader(HeaderAuthorizationTokenFrom(c)).
		Get(c.PocketBase.URL + EndpointLogs)

	if err != nil {
		return nil, err
	}

	switch res.StatusCode() {
	case 400:
		return nil, Error("something went wrong while processing your request. invalid filter")
	case 401:
		return nil, Error("the request requires admin authorization token to be set")
	case 403:
		return nil, Error("not allowed to perform request")
	}

	logs := Pagination[Log]{}
	if err := Unmarshal(res.Body(), &logs); err != nil {
		return nil, err
	}

	return &logs, nil
}

func (c *PocketClient) GetLogByID(id string) (*Log, error) {
	if !c.IsAuthenticated() {
		return nil, ErrNotAuthenticated
	}

	res, err := c.Resty.R().
		SetHeader(HeaderAuthorizationTokenFrom(c)).
		Get(c.PocketBase.URL + EndpointLogs + "/" + id)

	if err != nil {
		return nil, err
	}

	switch res.StatusCode() {
	case 401:
		return nil, Error("the request requires admin authorization token to be set")
	case 403:
		return nil, Error("not allowed to perform request")
	case 404:
		return nil, Error("log not found with provided id (" + id + ")")
	}

	log := Log{}
	if err := Unmarshal(res.Body(), &log); err != nil {
		return nil, err
	}

	return &log, nil
}

type LogStat struct {
	Total int    `json:"total"`
	Date  string `json:"date"`
}

func (c *PocketClient) GetLogggingStats() (*[]LogStat, error) {
	if !c.IsAuthenticated() {
		return nil, ErrNotAuthenticated
	}

	res, err := c.Resty.R().
		SetHeader(HeaderAuthorizationTokenFrom(c)).
		Get(c.PocketBase.URL + EndpointLogsStats)

	if err != nil {
		return nil, err
	}

	switch res.StatusCode() {
	case 400:
		return nil, Error("something went wrong while processing your request. invalid filter")
	case 401:
		return nil, Error("the request requires admin authorization token to be set")
	case 403:
		return nil, Error("not allowed to perform request")
	}

	stats := []LogStat{}
	if err := Unmarshal(res.Body(), &stats); err != nil {
		return nil, err
	}

	return &stats, nil
}
