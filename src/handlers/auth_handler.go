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
	ports "github.com/kiramishima/shining_guardian/src/ports/service"
	"github.com/unrolled/render"
	"go.uber.org/zap"
)

func NewAuthHandler(loggger *zap.SugaredLogger, svc ports.IAuthService, render *render.Render) *chi.Mux {
	var r = chi.NewRouter()
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.Compress(5))

	var h = &AuthHandler{
		logger:   loggger,
		service:  svc,
		response: render,
	}

	r.Route("/auth", func(r chi.Router) {
		r.Get("/heatlh", h.HealthHandler)
		r.Post("/sign-in", h.SignInHandler)
		r.Post("/sign-up", h.SignUpHandler)
	})

	return r
}

type AuthHandler struct {
	logger   *zap.SugaredLogger
	service  ports.IAuthService
	response *render.Render
}

func (h *AuthHandler) HealthHandler(w http.ResponseWriter, req *http.Request) {
	if err := h.response.JSON(w, http.StatusAccepted, map[string]string{"status": "OK", "version": "1.0"}); err != nil {
		h.logger.Error(err)
		_ = h.response.JSON(w, http.StatusInternalServerError, map[string]string{"error": httpErrors.InternalServerError.Error()})
		return
	}
}

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

	resp, err := h.service.SignIn(ctx, form)
	if err != nil {
		h.logger.Error(err.Error())

		select {
		case <-ctx.Done():
			_ = h.response.JSON(w, http.StatusGatewayTimeout, httpErrors.ErrTimeout)
			// http.Error(w, httpErrors.ErrTimeout.Error(), http.StatusGatewayTimeout)
		default:
			if errors.Is(err, httpErrors.ErrInvalidRequestBody) {
				_ = h.response.JSON(w, http.StatusBadRequest, httpErrors.ErrBadEmailOrPassword)
				// http.Error(w, httpErrors.ErrBadEmailOrPassword.Error(), http.StatusBadRequest)
			} else {
				_ = h.response.JSON(w, http.StatusInternalServerError, httpErrors.ErrBadEmailOrPassword)
				// http.Error(w, httpErrors.ErrBadEmailOrPassword.Error(), http.StatusInternalServerError)
			}
		}
		return
	}

	if err := h.response.JSON(w, http.StatusAccepted, resp); err != nil {
		h.logger.Error(err)
		_ = h.response.JSON(w, http.StatusInternalServerError, map[string]string{"error": httpErrors.InternalServerError.Error()})
		// http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

}

func (h *AuthHandler) SignUpHandler(w http.ResponseWriter, req *http.Request) {
	if err := h.response.JSON(w, http.StatusAccepted, map[string]string{"status": "OK", "version": "1.0"}); err != nil {
		h.logger.Error(err)
		_ = h.response.JSON(w, http.StatusInternalServerError, map[string]string{"error": httpErrors.InternalServerError.Error()})
		return
	}
}
