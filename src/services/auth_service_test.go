package services

import (
	"context"
	"github.com/kiramishima/shining_guardian/mocks"
	"github.com/kiramishima/shining_guardian/src/domain"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestAuthService_FindByCredentials(t *testing.T) {
	logger, _ := zap.NewProduction()
	slogger := logger.Sugar()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repo := mocks.NewMockIAuthRepository(mockCtrl)
	repo.EXPECT().FindByCredentials(gomock.Any(), gomock.Any()).Return(&domain.User{
		ID:        "1",
		Email:     "gini@mail.com",
		UserName:  "",
		Password:  "12356",
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}, nil)

	uc := NewAuthService(repo, slogger)

	t.Run("OK", func(t *testing.T) {
		ctx := context.Background()
		data := &domain.AuthRequest{Email: "gini@mail.com", Password: "123456"}
		b, err := uc.FindByCredentials(ctx, data)
		t.Log("B", b)
		assert.NoError(t, err)
		assert.Equal(t, "gini@mail.com", b.Token)
	})
	t.Run("Not Found", func(t *testing.T) {
		ctx := context.Background()
		data := &domain.AuthRequest{Email: "gini@mail.com", Password: ""}
		b, err := uc.FindByCredentials(ctx, data)
		t.Log(b)
		t.Log(err)
		assert.Error(t, err)
	})
}

func TestAuthService_Register(t *testing.T) {

}
