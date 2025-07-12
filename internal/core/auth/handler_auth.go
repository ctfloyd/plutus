package user

import (
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_handler"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_logger"
	"github.com/go-chi/chi/v5"
	chiWare "github.com/go-chi/chi/v5/middleware"
	"net/http"
	"plutus/internal/common/auth"
	"plutus/internal/common/handler"
	"plutus/internal/common/service_error"
	"plutus/pkg/plutus"
)

type Handler struct {
	logger     hz_logger.Logger
	authorizer *auth.Authorizer
	service    Service
}

func NewHandler(logger hz_logger.Logger, authorizer *auth.Authorizer, service Service) *Handler {
	return &Handler{logger: logger, authorizer: authorizer, service: service}
}

func (h *Handler) RegisterRoutes(mux *chi.Mux, hctx handler.Context) {
	if hctx.Version == handler.ApiVersionV1 {
		mux.Route("/v1/auth", func(r chi.Router) {
			r.Use(chiWare.Timeout(hctx.Timeout))
			r.Post("/login", h.Login)
			r.Post("/signup", h.Signup)
			r.Post("/refresh", h.Refresh)
		})
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request plutus.LoginRequest
	if ok := handler.ReadBody(w, r, &request); !ok {
		return
	}

	response, err := h.service.Login(ctx, request)
	if err != nil {
		h.logger.ErrorArgs(ctx, "An error occurred while logging a user in . %+v", err)
		handler.Error(ctx, w, service_error.Internal, "Failed to login.")
		return
	}

	hz_handler.Ok(w, response)

}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request plutus.SignupRequest
	if ok := handler.ReadBody(w, r, &request); !ok {
		return
	}

	response, err := h.service.Signup(ctx, request)
	if err != nil {
		h.logger.ErrorArgs(ctx, "An error occurred while signing up a user. %+v", err)
		handler.Error(ctx, w, service_error.Internal, "Failed to signup user.")
		return
	}

	hz_handler.Ok(w, response)
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request plutus.RefreshRequest
	if ok := handler.ReadBody(w, r, &request); !ok {
		return
	}

	response, err := h.service.Refresh(ctx, request)
	if err != nil {
		h.logger.ErrorArgs(ctx, "An error occurred while refreshing tokens. %+v", err)
		handler.Error(ctx, w, service_error.Internal, "Failed to refresh tokens.")
		return
	}

	hz_handler.Ok(w, response)
}
