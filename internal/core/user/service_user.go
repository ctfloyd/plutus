package user

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_logger"
	"github.com/oklog/ulid/v2"
	"plutus/internal/common/meta"
	"plutus/internal/common/transaction"
	"plutus/internal/gen/db"
	"plutus/pkg/plutus"
	"time"
)

var (
	ErrEmailInUse = errors.New("email is already in use")
)

type Service interface {
	GetUserById(ctx context.Context, id string) (plutus.GetUserByIdResponse, error)
	GetUserByEmail(ctx context.Context, email string) (plutus.GetUserByEmailResponse, error)
	CreateUser(ctx context.Context, request plutus.CreateUserRequest) (plutus.CreateUserResponse, error)
	IsEmailTaken(ctx context.Context, email string) (bool, error)
}

type DefaultService struct {
	logger     hz_logger.Logger
	txMgr      *transaction.Manager
	repository Repository
}

func NewService(logger hz_logger.Logger, txMgr *transaction.Manager, repository Repository) *DefaultService {
	return &DefaultService{logger: logger, txMgr: txMgr, repository: repository}
}

func (s *DefaultService) GetUserByEmail(ctx context.Context, email string) (plutus.GetUserByEmailResponse, error) {
	var user db.User
	var err error
	f := func(ctx context.Context) error {
		user, err = s.repository.GetUserByEmail(ctx, email)
		return err
	}

	if err := s.txMgr.WithTransaction(ctx, f); err != nil {
		return plutus.GetUserByEmailResponse{}, errors.WithStack(err)
	}

	return plutus.GetUserByEmailResponse{
		User: s.convertToApi(user),
		Meta: meta.Generate(ctx),
	}, nil
}

func (s *DefaultService) IsEmailTaken(ctx context.Context, email string) (bool, error) {
	var taken bool
	var err error
	f := func(ctx context.Context) error {
		taken, err = s.repository.IsEmailTaken(ctx, email)
		return err
	}

	if err := s.txMgr.WithTransaction(ctx, f); err != nil {
		return true, errors.WithStack(err)
	}

	return taken, err
}

func (s *DefaultService) GetUserById(ctx context.Context, id string) (plutus.GetUserByIdResponse, error) {
	var user db.User
	var err error
	f := func(ctx context.Context) error {
		user, err = s.repository.GetUserById(ctx, id)
		return err
	}

	if err := s.txMgr.WithTransaction(ctx, f); err != nil {
		return plutus.GetUserByIdResponse{}, errors.WithStack(err)
	}

	return plutus.GetUserByIdResponse{
		User: s.convertToApi(user),
		Meta: meta.Generate(ctx),
	}, nil
}

func (s *DefaultService) CreateUser(ctx context.Context, request plutus.CreateUserRequest) (plutus.CreateUserResponse, error) {
	var user db.User
	var err error
	f := func(ctx context.Context) error {
		taken, err := s.repository.IsEmailTaken(ctx, request.Email)
		if err != nil {
			return err
		}

		if taken {
			return ErrEmailInUse
		}

		params := db.CreateUserParams{
			ID:        ulid.Make().String(),
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Email:     request.Email,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		user, err = s.repository.InsertUser(ctx, params)
		return err
	}

	if err = s.txMgr.WithTransaction(ctx, f); err != nil {
		return plutus.CreateUserResponse{}, errors.WithStack(err)
	}

	return plutus.CreateUserResponse{
		User: s.convertToApi(user),
		Meta: meta.Generate(ctx),
	}, nil
}

func (s *DefaultService) convertToApi(user db.User) plutus.User {
	return plutus.User{
		Id:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
