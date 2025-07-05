package handler

import (
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_handler"
	"github.com/go-chi/chi/v5"
	"net/http"
	"plutus/internal/middleware"
	"plutus/internal/service_error"
	"time"
)

type Context struct {
	Timeout    time.Duration
	Authorizer *middleware.Authorizer
	Version    ApiVersion
}

type ApiVersion int

const (
	_ ApiVersion = iota
	ApiVersionV1
)

type PlutusHandler interface {
	RegisterRoutes(mux *chi.Mux, hctx Context)
}

func Unauthorized(w http.ResponseWriter) {
	hz_handler.Error(w, service_error.Unauthorized, "You are not authorized to access this resource.")
}
