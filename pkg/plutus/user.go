package plutus

import (
	"time"
)

type GetUserByIdResponse struct {
	User User `json:"user"`
	Meta Meta `json:"meta"`
}

type CreateUserRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type User struct {
	Id        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserResponse struct {
	User User `json:"user"`
	Meta Meta `json:"meta"`
}
