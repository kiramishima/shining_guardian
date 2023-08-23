package repository

import (
	"context"

	"github.com/kiramishima/shining_guardian/src/domain"
)

type IAuthRepository interface {
	FindByCredentials(ctx context.Context, data *domain.AuthRequest) (*domain.User, error)
	Register(ctx context.Context, registerReq *domain.RegisterRequest) error
}
