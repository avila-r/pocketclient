package pocketclient

type AdminProfile struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created"`
	UpdatedAt string `json:"updated"`
	Email     string `json:"email"`
	Avatar    int    `json:"avatar"`
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

func (m *ModuleAdmin) GetByID(id string) {
}

func (m *ModuleAdmin) New(email, password string) {
}

func (m *ModuleAdmin) Update(id struct{}) {
}

func (m *ModuleAdmin) DeleteByID(id struct{}) {

}
