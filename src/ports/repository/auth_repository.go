package repository

import (
	"context"

	"github.com/kiramishima/shining_guardian/src/domain"
)

type IAuthRepository interface {
	SignIn(ctx context.Context, data domain.AuthRequest) (*domain.User, error)
}
