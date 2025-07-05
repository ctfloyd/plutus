package user

import (
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_handler"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_logger"
	"github.com/go-chi/chi/v5"
	chiWare "github.com/go-chi/chi/v5/middleware"
	"net/http"
	"plutus/internal/auth"
	"plutus/internal/handler"
	"plutus/internal/service_error"
)

type Handler struct {
	logger  hz_logger.Logger
	service Service
}

func NewHandler(logger hz_logger.Logger, service Service) *Handler {
	return &Handler{logger: logger, service: service}
}

func (uh *Handler) RegisterRoutes(mux *chi.Mux, hctx handler.Context) {
	if hctx.Version == handler.ApiVersionV1 {
		mux.Route("/v1/user", func(r chi.Router) {
			r.Use(chiWare.Timeout(hctx.Timeout))
			r.Get("/{id}", uh.GetUserById)
		})
	}
}

func (uh *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	// TODO (c.floyd): Pass around an authorizer object that can be disabled.
	if !auth.IsMeOrRole(ctx, id, auth.RoleAdmin) {
		handler.Unauthorized(w)
		return
	}

	user, err := uh.service.GetUserById(ctx, id)
	if err != nil {
		hz_handler.Error(w, service_error.Internal, "Could not get user by id.")
		return
	}

	hz_handler.Ok(w, user)
}
