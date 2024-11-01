package pocketclient

type AdminCredentials struct {
	Token    string
	ID       string
	Email    string
	Password string
}

func (m *ModuleAdmin) Auth(email, password string) (*AdminCredentials, error) {
	url := Client.PocketBase.URL + EndpointAdminsAuthWithPassword

	request := Json{
		"identity": email,
		"password": password,
	}

	target := struct {
		Token        string `json:"token"`
		AdminDetails struct {
			ID     string `json:"id"`
			Email  string `json:"email"`
			Avatar int    `json:"avatar"`
		} `json:"admin"`
	}{}

	response, err := Client.Resty.R().
		SetBody(request).
		Post(url)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() == 400 {
		return nil, Error("unable to submit form - " + string(response.Body()))
	}

	if err := Unmarshal(response.Body(), &target); err != nil {
		return nil, err
	}

	credentials := AdminCredentials{
		Token:    target.Token,
		ID:       target.AdminDetails.ID,
		Email:    target.AdminDetails.Email,
		Password: password,
	}

	return &credentials, nil
}
