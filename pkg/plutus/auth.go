package plutus

type SignupRequest struct {
	User     CreateUserRequest `json:"user"`
	Password string            `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token TokenResponse `json:"token"`
	Meta  Meta          `json:"meta"`
}

type TokenResponse struct {
	AuthorizationToken string `json:"authorization_token"`
	RefreshToken       string `json:"refresh_token"`
}

type SignupResponse struct {
	Token TokenResponse `json:"token"`
	Meta  Meta          `json:"meta"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	Token TokenResponse `json:"token"`
	Meta  Meta          `json:"meta"`
}
