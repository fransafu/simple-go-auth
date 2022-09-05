package handlers

import "net/http"

type IAuthHandler interface {
	Login(http.ResponseWriter, *http.Request)
	Register(http.ResponseWriter, *http.Request)
	Logout(http.ResponseWriter, *http.Request)
}

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// valid params

	// call service

	// response
}

func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	// valid params

	// call service

	// response
}

func (a *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// valid params

	// call service

	// response
}
