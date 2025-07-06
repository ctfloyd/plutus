package auth

import (
	"context"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_logger"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	ContextKeyToken  = "x-plutus-token"
	ContextKeyRoles  = "x-plutus-roles"
	ContextKeyUserId = "x-plutus-user-id"
)

type Role = string

const (
	RoleAdmin Role = "x-plutus-admin"
)

type PlutusClaims struct {
	jwt.RegisteredClaims
	Roles jwt.ClaimStrings `json:"roles,omitempty"`
}

type Authorizer struct {
	logger      hz_logger.Logger
	enforceAuth bool
	secret      []byte
}

func NewAuthorizer(logger hz_logger.Logger, enforceAuth bool, secret string) *Authorizer {
	return &Authorizer{logger: logger, enforceAuth: enforceAuth, secret: []byte(secret)}
}

func (a *Authorizer) GenerateJWT(userId string, roles []string) (string, error) {
	return jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		&PlutusClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "plutus",
				Subject:   userId,
				Audience:  []string{"user"},
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
			Roles: roles,
		}).SignedString(a.secret)
}

func (a *Authorizer) IsMe(ctx context.Context, id string) bool {
	if !a.enforceAuth {
		return true
	}

	if userId, ok := ctx.Value(ContextKeyUserId).(string); ok {
		return userId == id
	}

	return false
}

func (a *Authorizer) HasRole(ctx context.Context, role Role) bool {
	if !a.enforceAuth {
		return true
	}

	if roles, ok := ctx.Value(ContextKeyRoles).(jwt.ClaimStrings); ok {
		for _, r := range roles {
			if r == role {
				return true
			}
		}
	}
	return false
}

func (a *Authorizer) IsMeOrRole(ctx context.Context, meId string, role string) bool {
	if !a.enforceAuth {
		return true
	}

	return a.IsMe(ctx, meId) || a.HasRole(ctx, role)
}
