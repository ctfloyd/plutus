package user

import (
	"github.com/cockroachdb/errors"
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

func (uh *Handler) RegisterRoutes(mux *chi.Mux, hctx handler.Context) {
	if hctx.Version == handler.ApiVersionV1 {
		mux.Route("/v1/user", func(r chi.Router) {
			r.Use(chiWare.Timeout(hctx.Timeout))
			r.Get("/{id}", uh.GetUserById)
			r.Post("/", uh.CreateUser)
		})
	}
}

func (uh *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if !uh.authorizer.HasRole(ctx, auth.RoleAdmin) {
		handler.Unauthorized(ctx, w)
		return
	}

	var request plutus.CreateUserRequest
	if ok := handler.ReadBody(w, r, &request); !ok {
		return
	}

	response, err := uh.service.CreateUser(ctx, request)
	if err != nil {
		if errors.Is(err, ErrEmailInUse) {
			uh.logger.Warn(ctx, "Email is already in use!")
			handler.Error(ctx, w, service_error.BadRequest, "Failed to create user.")
		} else {
			uh.logger.ErrorArgs(ctx, "An error occurred while creating a user. %+v", err)
			handler.Error(ctx, w, service_error.Internal, "Failed to create user.")
		}
		return
	}

	hz_handler.Ok(w, response)
}

func (uh *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	if !uh.authorizer.IsMeOrRole(ctx, id, auth.RoleAdmin) {
		handler.Unauthorized(ctx, w)
		return
	}

	user, err := uh.service.GetUserById(ctx, id)
	if err != nil {
		handler.Error(ctx, w, service_error.Internal, "Could not get user by id.")
		return
	}

	hz_handler.Ok(w, user)
}
