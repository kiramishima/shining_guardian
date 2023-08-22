package services

import (
	"context"

	"github.com/kiramishima/shining_guardian/src/domain"
	ports "github.com/kiramishima/shining_guardian/src/ports/repository"
)

type AuthService struct {
	repository ports.IAuthRepository
}

func NewAuthService(repo ports.IAuthRepository) *AuthService {
	return &AuthService{
		repository: repo,
	}
}

func (svc *AuthService) SignIn(ctx context.Context, data *domain.AuthRequest) (*domain.AuthResponse, error) {
	return nil, nil
}
