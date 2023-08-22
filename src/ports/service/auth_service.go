package service

import (
	"context"

	"github.com/kiramishima/shining_guardian/src/domain"
)

type IAuthService interface {
	SignIn(ctx context.Context, data *domain.AuthRequest) (*domain.AuthResponse, error)
}
