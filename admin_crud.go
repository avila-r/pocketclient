package pocketclient

import "github.com/avila-r/pocketclient/validation"

type AdminProfile struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created"`
	UpdatedAt string `json:"updated"`
	Email     string `json:"email"`
	Avatar    int    `json:"avatar"`
}

type AdminRequest struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   int    `json:"avatar"`
}

func (m *ModuleAdmin) ListAll(p ...PaginationParams) (*Pagination[AdminProfile], error) {
	params := PaginationParams{}
	if len(p) > 0 {
		params = p[0]
	}

	if !Client.IsAuthenticated() {
		return nil, ErrNotAuthenticated
	}

	res, err := Client.Resty.R().
		SetQueryParams(params.ToQueryParams()).
		SetHeader(HeaderAuthorizationToken()).
		Get(Client.PocketBase.URL + EndpointAdmins)

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

	admins := Pagination[AdminProfile]{}
	if err := Unmarshal(res.Body(), &admins); err != nil {
		return nil, err
	}

	return &admins, nil
}

func (m *ModuleAdmin) GetByID(id string) (*AdminProfile, error) {
	if !Client.IsAuthenticated() {
		return nil, ErrNotAuthenticated
	}

	res, err := Client.Resty.R().
		SetHeader(HeaderAuthorizationToken()).
		Get(Client.Resty.BaseURL + EndpointAdmins + "/" + id)

	if err != nil {
		return nil, err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return nil, err
	}

	admin := AdminProfile{}
	if err := Unmarshal(res.Body(), &admin); err != nil {
		return nil, err
	}

	return &admin, nil
}

func (m *ModuleAdmin) New(new AdminRequest) (*AdminProfile, error) {
	if !Client.IsAuthenticated() {
		return nil, ErrNotAuthenticated
	}

	res, err := Client.Resty.R().
		SetHeader(HeaderAuthorizationToken()).
		SetBody(new).
		Post(Client.PocketBase.URL + EndpointAdmins)

	if err != nil {
		return nil, err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return nil, err
	}

	created := AdminProfile{}
	if err := Unmarshal(res.Body(), &created); err != nil {
		return nil, err
	}

	return &created, nil
}

func (m *ModuleAdmin) Update(id string, new AdminRequest) (*AdminProfile, error) {
	if !Client.IsAuthenticated() {
		return nil, ErrNotAuthenticated
	}

	new.ID = ""

	res, err := Client.Resty.R().
		SetHeader(HeaderAuthorizationToken()).
		SetBody(new).
		Patch(Client.PocketBase.URL + EndpointAdmins + "/" + id)

	if err != nil {
		return nil, err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return nil, err
	}

	updated := AdminProfile{}
	if err := Unmarshal(res.Body(), &updated); err != nil {
		return nil, err
	}

	return &updated, nil
}

func (m *ModuleAdmin) DeleteByID(id string) error {
	if !Client.IsAuthenticated() {
		return ErrNotAuthenticated
	}

	res, err := Client.Resty.R().
		SetHeader(HeaderAuthorizationToken()).
		Delete(Client.PocketBase.URL + EndpointAdmins + "/" + id)

	if err != nil {
		return err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return err
	}

	return nil
}
