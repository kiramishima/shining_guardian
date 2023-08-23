package services

import (
	"context"
	"errors"
	"fmt"
	httpErrors "github.com/kiramishima/shining_guardian/pkg/errors"

	"github.com/golang-jwt/jwt"
	"github.com/kiramishima/shining_guardian/pkg/utils/jwtutils"
	"github.com/kiramishima/shining_guardian/src/domain"
	ports "github.com/kiramishima/shining_guardian/src/ports/repository"
	port "github.com/kiramishima/shining_guardian/src/ports/service"
	"go.uber.org/zap"
)

type AuthService struct {
	logger     *zap.SugaredLogger
	repository ports.IAuthRepository
}

var _ port.IAuthService = (*AuthService)(nil)

func NewAuthService(repo ports.IAuthRepository, logger *zap.SugaredLogger) *AuthService {
	return &AuthService{
		logger:     logger,
		repository: repo,
	}
}

// FindByCredentials service method for retrieve user if provide credentials are valid
func (svc *AuthService) FindByCredentials(ctx context.Context, data *domain.AuthRequest) (*domain.AuthResponse, error) {
	data.Password = data.Hash256Password(data.Password)
	user, err := svc.repository.FindByCredentials(ctx, data)

	if err != nil {
		svc.logger.Error(err.Error())

		select {
		case <-ctx.Done():
			return nil, httpErrors.ErrTimeout
		default:
			if errors.Is(err, httpErrors.ErrInvalidRequestBody) {
				return nil, httpErrors.ErrBadEmailOrPassword
			} else {
				return nil, httpErrors.ErrBadEmailOrPassword
			}
		}
	}

	token, err := jwtutils.GenerateJWT(user)
	if err != nil {
		svc.logger.Error(err.Error(), fmt.Sprintf("%T", err))
		return nil, jwt.ErrSignatureInvalid
	}

	return &domain.AuthResponse{Token: token}, nil

}

// Register repository method for create a new user.
func (svc *AuthService) Register(ctx context.Context, registerReq *domain.RegisterRequest) error {
	registerReq.Password = registerReq.Hash256Password(registerReq.Password)
	err := svc.repository.Register(ctx, registerReq)

	if err != nil {
		svc.logger.Error(err.Error())

		select {
		case <-ctx.Done():
			return httpErrors.ErrTimeout
		default:
			if errors.Is(err, httpErrors.ErrInvalidRequestBody) {
				return httpErrors.ErrBadEmailOrPassword
			} else {
				return httpErrors.ErrBadEmailOrPassword
			}
		}
	}

	return nil
}
