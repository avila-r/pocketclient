package pocketclient

import "github.com/avila-r/pocketclient/validation"

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

	if err := validation.VerifyResponse(res); err != nil {
		return nil, err
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

	if err := validation.VerifyResponse(res); err != nil {
		return nil, err
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

	if err := validation.VerifyResponse(res); err != nil {
		return nil, err
	}

	stats := []LogStat{}
	if err := Unmarshal(res.Body(), &stats); err != nil {
		return nil, err
	}

	return &stats, nil
}
