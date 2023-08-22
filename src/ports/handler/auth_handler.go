package handler

import "net/http"

type IAuthHandler interface {
	HealthHandler(w http.ResponseWriter, req *http.Request)
	SignInHandler(w http.ResponseWriter, req *http.Request)
	SignUpHandler(w http.ResponseWriter, req *http.Request)
}
