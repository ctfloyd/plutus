package user

import (
	"context"
	"crypto/rand"
	"github.com/cockroachdb/errors"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_logger"
	"golang.org/x/crypto/bcrypt"
	"plutus/internal/common/auth"
	"plutus/internal/common/meta"
	"plutus/internal/common/transaction"
	"plutus/internal/core/user"
	"plutus/internal/gen/db"
	"plutus/pkg/plutus"
	"time"
)

var (
	ErrEmailInUse = errors.New("email is already in use")
)

type Service interface {
	Login(ctx context.Context, request plutus.LoginRequest) (plutus.LoginResponse, error)
	Signup(ctx context.Context, request plutus.SignupRequest) (plutus.SignupResponse, error)
	Refresh(ctx context.Context, request plutus.RefreshRequest) (plutus.RefreshResponse, error)
}

type DefaultService struct {
	logger      hz_logger.Logger
	txMgr       *transaction.Manager
	repository  Repository
	userService user.Service
	authorizer  *auth.Authorizer
	pwCost      int
	saltSize    int
}

func NewService(
	logger hz_logger.Logger,
	txMgr *transaction.Manager,
	repository Repository,
	userService user.Service,
	authorizer *auth.Authorizer,
	pwCost int,
	saltSize int,
) *DefaultService {
	return &DefaultService{
		logger:      logger,
		txMgr:       txMgr,
		repository:  repository,
		userService: userService,
		authorizer:  authorizer,
		pwCost:      pwCost,
		saltSize:    saltSize,
	}
}

func (s *DefaultService) Refresh(ctx context.Context, request plutus.RefreshRequest) (plutus.RefreshResponse, error) {
	var err error
	var tokens tokens

	f := func(ctx context.Context) error {
		var token db.Token
		if token, err = s.repository.GetToken(ctx, request.RefreshToken); err != nil {
			return errors.WithStack(err)
		}

		if token.Revoked {
			return errors.New("token is revoked")
		}

		if tokens, err = s.generateTokens(token.UserID, []string{auth.RoleUser}); err != nil {
			return errors.WithStack(err)
		}

		refreshParams := db.CreateTokenParams{
			UserID:  token.UserID,
			Token:   tokens.Refresh,
			Revoked: false,
		}
		if err = s.repository.InsertToken(ctx, refreshParams); err != nil {
			return errors.WithStack(err)
		}

		if err = s.repository.RevokeToken(ctx, token.Token); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}

	if err = s.txMgr.WithTransaction(ctx, f); err != nil {
		return plutus.RefreshResponse{}, errors.WithStack(err)
	}

	return plutus.RefreshResponse{
		Token: plutus.TokenResponse{
			AuthorizationToken: tokens.Auth,
			RefreshToken:       tokens.Refresh,
		},
		Meta: meta.Generate(ctx),
	}, nil
}

func (s *DefaultService) Login(ctx context.Context, request plutus.LoginRequest) (plutus.LoginResponse, error) {
	var err error
	var tokens tokens

	f := func(ctx context.Context) error {
		var userId string
		if userResponse, err := s.userService.GetUserByEmail(ctx, request.Email); err != nil {
			return errors.WithStack(err)
		} else {
			userId = userResponse.User.Id
		}

		var userAuth db.Auth
		if userAuth, err = s.repository.GetAuth(ctx, userId); err != nil {
			return errors.WithStack(err)
		}

		if verified := s.verifyPassword(request.Password, userAuth.PasswordHash, userAuth.Salt); !verified {
			return errors.WithStack(errors.New("password is incorrect"))
		}

		if tokens, err = s.generateTokens(userId, []string{auth.RoleUser}); err != nil {
			return errors.WithStack(err)
		}

		refreshParams := db.CreateTokenParams{
			UserID:  userId,
			Token:   tokens.Refresh,
			Revoked: false,
		}
		if err = s.repository.InsertToken(ctx, refreshParams); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}

	if err = s.txMgr.WithTransaction(ctx, f); err != nil {
		return plutus.LoginResponse{}, errors.WithStack(err)
	}

	return plutus.LoginResponse{
		Token: plutus.TokenResponse{
			AuthorizationToken: tokens.Auth,
			RefreshToken:       tokens.Refresh,
		},
		Meta: meta.Generate(ctx),
	}, nil
}

func (s *DefaultService) Signup(ctx context.Context, request plutus.SignupRequest) (plutus.SignupResponse, error) {
	var err error
	var tokens tokens
	f := func(ctx context.Context) error {
		if taken, err := s.userService.IsEmailTaken(ctx, request.User.Email); err != nil {
			return errors.WithStack(err)
		} else if taken {
			return ErrEmailInUse
		}

		createUserParams := plutus.CreateUserRequest{
			FirstName: request.User.FirstName,
			LastName:  request.User.LastName,
			Email:     request.User.Email,
		}
		var userId string
		if response, err := s.userService.CreateUser(ctx, createUserParams); err != nil {
			return errors.WithStack(err)
		} else {
			userId = response.User.Id
		}

		if err = s.createAuth(ctx, userId, request.Password); err != nil {
			return errors.WithStack(err)
		}

		if tokens, err = s.generateTokens(userId, []string{auth.RoleUser}); err != nil {
			return errors.WithStack(err)
		}

		refreshParams := db.CreateTokenParams{
			UserID:  userId,
			Token:   tokens.Refresh,
			Revoked: false,
		}
		if err = s.repository.InsertToken(ctx, refreshParams); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}

	if err = s.txMgr.WithTransaction(ctx, f); err != nil {
		return plutus.SignupResponse{}, errors.WithStack(err)
	}

	return plutus.SignupResponse{
		Token: plutus.TokenResponse{
			AuthorizationToken: tokens.Auth,
			RefreshToken:       tokens.Refresh,
		},
		Meta: meta.Generate(ctx),
	}, nil

}

func (s *DefaultService) createAuth(ctx context.Context, userId string, password string) error {
	salt, err := s.generateSalt()
	if err != nil {
		return errors.WithStack(err)
	}

	hashResult, err := s.hashPasswordWithSalt(password, salt)
	if err != nil {
		return errors.WithStack(err)
	}

	createAuthParams := db.CreateAuthParams{
		UserID:       userId,
		PasswordHash: hashResult.Hash,
		Salt:         hashResult.Salt,
		CreatedAt:    time.Now(),
	}
	err = s.repository.InsertAuth(ctx, createAuthParams)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
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

type tokens struct {
	Auth    string
	Refresh string
}

func (s *DefaultService) generateTokens(userId string, roles []string) (tokens, error) {
	authToken, err := s.authorizer.GenerateAuthenticationToken(userId, roles)
	if err != nil {
		return tokens{}, err
	}
	refreshToken, err := s.authorizer.GenerateRefreshToken(userId)
	if err != nil {
		return tokens{}, err
	}

	return tokens{
		Auth:    authToken,
		Refresh: refreshToken,
	}, nil

}

func (s *DefaultService) generateSalt() ([]byte, error) {
	salt := make([]byte, s.saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

type HashResult struct {
	Hash []byte
	Salt []byte
}

func (s *DefaultService) hashPassword(password string) (HashResult, error) {
	salt, err := s.generateSalt()
	if err != nil {
		return HashResult{}, err
	}
	return s.hashPasswordWithSalt(password, salt)
}

func (s *DefaultService) hashPasswordWithSalt(password string, salt []byte) (HashResult, error) {
	saltedPassword := append([]byte(password), salt...)

	hash, err := bcrypt.GenerateFromPassword(saltedPassword, s.pwCost)
	if err != nil {
		return HashResult{}, err
	}

	return HashResult{
		Hash: hash,
		Salt: salt,
	}, nil
}

func (s *DefaultService) verifyPassword(password string, hash []byte, salt []byte) bool {
	// Combine password and salt
	saltedPassword := append([]byte(password), salt...)

	err := bcrypt.CompareHashAndPassword(hash, saltedPassword)

	dummyHash := []byte("$2a$12$dummy.hash.to.prevent.timing.attacks")
	if err != nil {
		_ = bcrypt.CompareHashAndPassword(dummyHash, saltedPassword)
		return false
	}

	return true
}
