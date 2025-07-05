package auth

import "context"

const (
	ContextKeyToken  = "x-plutus-token"
	ContextKeyRoles  = "x-plutus-roles"
	ContextKeyUserId = "x-plutus-user-id"
)

type Role = string

const (
	RoleAdmin Role = "x-plutus-admin"
)

func IsMe(ctx context.Context, id string) bool {
	if userId, ok := ctx.Value(ContextKeyUserId).(string); ok {
		return userId == id
	}

	return false
}

func HasRole(ctx context.Context, role Role) bool {
	if roles, ok := ctx.Value(ContextKeyRoles).([]string); ok {
		for _, r := range roles {
			if r == role {
				return true
			}
		}
	}
	return false
}

func IsMeOrRole(ctx context.Context, meId string, role string) bool {
	return IsMe(ctx, meId) || HasRole(ctx, role)
}
