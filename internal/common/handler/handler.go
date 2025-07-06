package handler

import (
	"context"
	"fmt"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_api"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_handler"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_service_error"
	"github.com/go-chi/chi/v5"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
	"plutus/internal/common/meta"
	"plutus/internal/common/service_error"
	"plutus/pkg/plutus"
	"time"
)

type Context struct {
	Timeout time.Duration
	Version ApiVersion
}

type ApiVersion int

const (
	_ ApiVersion = iota
	ApiVersionV1
)

type PlutusHandler interface {
	RegisterRoutes(mux *chi.Mux, hctx Context)
}

func Error(ctx context.Context, w http.ResponseWriter, e hz_service_error.ServiceError, msg string) {
	hz_handler.ErrorWithResponse(w, convertToResponseError(ctx, e, msg))
}

func Unauthorized(ctx context.Context, w http.ResponseWriter) {
	Error(ctx, w, service_error.Unauthorized, "You are not authorized to access this resource.")
}

func ErrorArgs(ctx context.Context, w http.ResponseWriter, e hz_service_error.ServiceError, msg string, args ...any) {
	Error(ctx, w, e, fmt.Sprintf(msg, args...))
}

func ReadBody(w http.ResponseWriter, r *http.Request, body any) bool {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		Error(r.Context(), w, hz_service_error.Internal, "An unexpected error occurred while reading request body.")
		return false
	}

	err = jsoniter.Unmarshal(bytes, body)
	if err != nil {
		Error(r.Context(), w, hz_service_error.BadRequest, fmt.Sprintf("The request body could not be parsed. %v", err))
		return false
	}

	return true
}

func convertToResponseError(ctx context.Context, serviceError hz_service_error.ServiceError, message string) hz_handler.ErrorResponse {
	return plutus.ErrorResponse{
		ErrorResponse: hz_api.ErrorResponse{
			Code:      serviceError.Code,
			Status:    serviceError.Status,
			Message:   message,
			Timestamp: time.Now(),
		},
		Meta: meta.Generate(ctx),
	}

}
