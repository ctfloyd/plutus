package user

import (
	"context"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_logger"
	"plutus/pkg/api"
)

type Service interface {
	GetUserById(ctx context.Context, id string) (api.GetUserByIdResponse, error)
}

type DefaultService struct {
	logger hz_logger.Logger
}

func NewService(logger hz_logger.Logger) *DefaultService {
	return &DefaultService{logger: logger}
}

func (s *DefaultService) GetUserById(ctx context.Context, id string) (api.GetUserByIdResponse, error) {
	return api.GetUserByIdResponse{}, nil
}
