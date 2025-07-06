package meta

import (
	"context"
	chiWare "github.com/go-chi/chi/v5/middleware"
	"plutus/pkg/plutus"
)

func Generate(ctx context.Context) plutus.Meta {
	return plutus.Meta{
		RequestId: ctx.Value(chiWare.RequestIDKey).(string),
	}
}
