package plutus

type SignupRequest struct {
	User     CreateUserRequest `json:"user"`
	Password string            `json:"password"`
}

type TokenResponse struct {
	AuthorizationToken string `json:"authorization_token"`
	RefreshToken       string `json:"refresh_token"`
}

type SignupResponse struct {
	Token TokenResponse `json:"token"`
}
