package service

import (
	"context"

	"github.com/kiramishima/shining_guardian/src/domain"
)

type IAuthService interface {
	FindByCredentials(ctx context.Context, data *domain.AuthRequest) (*domain.AuthResponse, error)
	Register(ctx context.Context, registerReq *domain.RegisterRequest) error
}
