package pocketclient

import "errors"

const (
	EndpointAdmins                = "/api/admins"
	EndpointAdminAuthWithPassword = EndpointAdmins + "/auth-with-password"
)

var (
	HeaderAuthorizationToken = func() (string, string) {
		return "Authorization", Client.PocketBase.FirstAdmin().Token
	}
)

var (
	ErrUnableToAuthorizeFirstAdmin = errors.New("unable to authorize first admin")
)
