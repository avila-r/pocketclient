package pocketclient

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
