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

func (m *ModuleAdmin) RequestPasswordReset(email string) error {
	url := Client.PocketBase.URL + EndpointAdminsRequestPasswordReset

	request := Json{
		"email": email,
	}

	response, err := Client.Resty.R().
		SetBody(request).
		Post(url)

	if err != nil {
		return err
	}

	if response.StatusCode() != 200 {
		return Error("unable to request password reset. check logs for detailed information")
	}

	return nil
}

type PasswordResetRequest struct {
	Token       string
	NewPassword string
}

func (m *ModuleAdmin) ConfirmPasswordReset(req PasswordResetRequest) error {
	url := Client.PocketBase.URL + EndpointAdminsConfirmPasswordReset

	request := Json{
		"token":           req.Token,
		"password":        req.NewPassword,
		"passwordConfirm": req.NewPassword,
	}

	response, err := Client.Resty.R().
		SetBody(request).
		Post(url)

	if err != nil {
		return err
	}

	if response.StatusCode() != 200 {
		return Error("unable to confirm password reset. check logs for detailed information")
	}

	return nil
}
