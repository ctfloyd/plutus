package middleware

import (
	"context"
	"errors"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_logger"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"plutus/internal/auth"
	"strings"
)

type AuthParser struct {
	logger hz_logger.Logger
	secret []byte
}

func NewAuthParser(logger hz_logger.Logger, secret string) *AuthParser {
	return &AuthParser{logger, []byte(secret)}
}

func (a *AuthParser) Parse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("Authorization")
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

		ctx := a.populateContextWithClaims(r.Context(), parts[1])
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (a *AuthParser) populateContextWithClaims(ctx context.Context, token string) context.Context {
	if token == "" {
		return ctx
	}

	ctx = context.WithValue(ctx, auth.ContextKeyToken, token)
	tok, err := jwt.ParseWithClaims(token, &auth.PlutusClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			a.logger.WarnArgs(ctx, "Unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New("unexpected signing method")
		}
		return a.secret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS512.Alg()}))

	if err != nil {
		return ctx
	}

	if claims, ok := tok.Claims.(*auth.PlutusClaims); ok && tok.Valid {
		if sub := claims.RegisteredClaims.Subject; sub != "" {
			ctx = context.WithValue(ctx, auth.ContextKeyUserId, sub)
		}

		if roles := claims.Roles; roles != nil && len(roles) > 0 {
			ctx = context.WithValue(ctx, auth.ContextKeyRoles, roles)
		}
	}

	return ctx
}
