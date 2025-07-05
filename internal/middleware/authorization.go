package middleware

import (
	"context"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_logger"
	"net/http"
	"plutus/internal/auth"
	"strings"
)

type Authorizer struct {
	logger hz_logger.Logger
}

func NewAuthorizer(logger hz_logger.Logger) *Authorizer {
	return &Authorizer{logger}
}

func (a *Authorizer) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("Authorization")

		ctx := context.WithValue(r.Context(), auth.ContextKeyToken, "")
		ctx = context.WithValue(ctx, auth.ContextKeyUserId, "userid")
		ctx = context.WithValue(ctx, auth.ContextKeyRoles, []string{auth.RoleAdmin})
		r = r.WithContext(ctx)

		if headerToken == "" {
			a.logger.Warn(r.Context(), "No authorization header.")
			next.ServeHTTP(w, r)
			return
		}

		parts := strings.Split(headerToken, " ")
		if len(parts) != 2 {
			a.logger.Warn(r.Context(), "Malformed authorization header.")
			next.ServeHTTP(w, r)
			return
		}

		// TODO (cfloyd): Make this middleware do the thing.
		//token := parts[1]

		next.ServeHTTP(w, r)
	})
}
