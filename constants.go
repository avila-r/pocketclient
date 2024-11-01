package pocketclient

const (
	EndpointAdmins                     = "/api/admins"
	EndpointAdminsAuthWithPassword     = EndpointAdmins + "/auth-with-password"
	EndpointAdminsRequestPasswordReset = EndpointAdmins + "/request-password-reset"
	EndpointAdminsConfirmPasswordReset = EndpointAdmins + "/confirm-password-reset"

	EndpointHealthCheck = "/api/health"

	EndpointLogs      = "/api/logs"
	EndpointLogsStats = EndpointLogs + "/stats"
)

var (
	HeaderAuthorizationToken = func() (string, string) {
		return "Authorization", Client.PocketBase.Credentials.Token
	}

	HeaderAuthorizationTokenFrom = func(c *PocketClient) (string, string) {
		return "Authorization", c.PocketBase.Credentials.Token
	}
)

var (
	ErrNotAuthenticated            = Error("pocketclient not authenticated (call [AdminModule] Admin.Login() to authenticate first admin)")
	ErrUnableToAuthorizeFirstAdmin = Error("unable to authorize first admin")
)
