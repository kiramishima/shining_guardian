package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/kiramishima/shining_guardian/mocks"
	"github.com/kiramishima/shining_guardian/src/domain"
	"github.com/stretchr/testify/assert"
	"github.com/unrolled/render"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthHandler_HealthHandler(t *testing.T) {

}

func TestAuthHandler_SignInHandler(t *testing.T) {
	testCases := map[string]struct {
		ID            any
		buildStubs    func(uc *mocks.MockIAuthService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		"OK": {
			ID: 1,
			buildStubs: func(uc *mocks.MockIAuthService) {
				uc.EXPECT().
					FindByCredentials(gomock.Any(), &domain.AuthRequest{Email: "gini@mail.com", Password: "123456"}).
					Times(1).
					Return(&domain.User{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusAccepted, recorder.Code)
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uc := mocks.NewMockIAuthService(ctrl)
			tc.buildStubs(uc)

			recorder := httptest.NewRecorder()

			url := "/auth/sign-in"
			teacher := domain.AuthRequest{
				Email:    "gini@mail.com",
				Password: "123456",
			}
			// marshall data to json (like json_encode)
			marshalled, err := json.Marshal(teacher)
			if err != nil {
				log.Fatalf("impossible to marshall form: %s", err)
			}

			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(marshalled))
			assert.NoError(t, err)

			// router := chi.NewRouter()
			logger, _ := zap.NewProduction()
			slogger := logger.Sugar()
			r := render.New()
			router := NewAuthHandler(slogger, uc, r)
			router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}

func TestAuthHandler_SignUpHandler(t *testing.T) {

}
