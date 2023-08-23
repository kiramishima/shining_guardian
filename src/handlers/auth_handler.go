package handlers

import (
	"errors"
	"net/http"
	"time"

	httpErrors "github.com/kiramishima/shining_guardian/pkg/errors"
	httpUtils "github.com/kiramishima/shining_guardian/pkg/utils/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kiramishima/shining_guardian/src/domain"
	port "github.com/kiramishima/shining_guardian/src/ports/handler"
	ports "github.com/kiramishima/shining_guardian/src/ports/service"
	"github.com/unrolled/render"
	"go.uber.org/zap"
)

var _ port.IAuthHandler = (*AuthHandler)(nil)

// NewAuthHandler creates a instance of chi Router with embeded routes
func NewAuthHandler(logger *zap.SugaredLogger, svc ports.IAuthService, render *render.Render) *chi.Mux {
	var r = chi.NewRouter()
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.Compress(5))

	var h = &AuthHandler{
		logger:   logger,
		service:  svc,
		response: render,
	}

	r.Route("/auth", func(r chi.Router) {
		r.Get("/health", h.HealthHandler)
		r.Post("/sign-in", h.SignInHandler)
		r.Post("/sign-up", h.SignUpHandler)
	})

	return r
}

// AuthHandler struct
type AuthHandler struct {
	logger   *zap.SugaredLogger
	service  ports.IAuthService
	response *render.Render
}

// HealthHandler for healthing check
func (h *AuthHandler) HealthHandler(w http.ResponseWriter, req *http.Request) {
	if err := h.response.JSON(w, http.StatusAccepted, map[string]string{"status": "OK", "version": "1.0"}); err != nil {
		h.logger.Error(err)
		_ = h.response.JSON(w, http.StatusInternalServerError, map[string]string{"error": httpErrors.InternalServerError.Error()})
		return
	}
}

// SignInHandler for sign users
func (h *AuthHandler) SignInHandler(w http.ResponseWriter, req *http.Request) {
	var form = &domain.AuthRequest{}

	err := httpUtils.ReadJSON(w, req, &form)

	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, httpErrors.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}
	h.logger.Info(form)
	ctx := req.Context()

	resp, err := h.service.FindByCredentials(ctx, form)
	if err != nil {
		h.logger.Error(err.Error())

		select {
		case <-ctx.Done():
			_ = h.response.JSON(w, http.StatusGatewayTimeout, httpErrors.ErrTimeout)
		default:
			if errors.Is(err, httpErrors.ErrInvalidRequestBody) {
				_ = h.response.JSON(w, http.StatusBadRequest, httpErrors.ErrBadEmailOrPassword)
			} else {
				_ = h.response.JSON(w, http.StatusInternalServerError, httpErrors.ErrBadEmailOrPassword)
			}
		}
		return
	}

	if err := h.response.JSON(w, http.StatusAccepted, resp); err != nil {
		h.logger.Error(err)
		_ = h.response.JSON(w, http.StatusInternalServerError, map[string]string{"error": httpErrors.InternalServerError.Error()})
		return
	}

}

// SignUpHandler for register new users
func (h *AuthHandler) SignUpHandler(w http.ResponseWriter, req *http.Request) {
	var form = &domain.RegisterRequest{}

	err := httpUtils.ReadJSON(w, req, &form)

	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, httpErrors.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}
	h.logger.Info(form)
	ctx := req.Context()

	err = h.service.Register(ctx, form)
	if err != nil {
		h.logger.Error(err.Error())

		select {
		case <-ctx.Done():
			_ = h.response.JSON(w, http.StatusGatewayTimeout, httpErrors.ErrTimeout)
		default:
			if errors.Is(err, httpErrors.ErrInvalidRequestBody) {
				_ = h.response.JSON(w, http.StatusBadRequest, httpErrors.ErrBadEmailOrPassword)
			} else {
				_ = h.response.JSON(w, http.StatusInternalServerError, httpErrors.ErrBadEmailOrPassword)
			}
		}
		return
	}

	if err := h.response.JSON(w, http.StatusAccepted, domain.SuccessResponse{Message: "Success. Check you email for activate your account."}); err != nil {
		h.logger.Error(err)
		_ = h.response.JSON(w, http.StatusInternalServerError, map[string]string{"error": httpErrors.InternalServerError.Error()})
		// http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}
