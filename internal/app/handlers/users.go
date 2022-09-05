package handlers

import (
	"net/http"

	"github.com/fransafu/simple-go-auth/internal/app/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type IUserHandler interface {
	GetAll(http.ResponseWriter, *http.Request)
	GetByID(http.ResponseWriter, *http.Request)
}

type UserHandler struct {
	userService services.IUserService
}

func NewUserHandler(userService services.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (a *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := a.userService.GetAll()
	if err != nil {
		render.Render(w, r, ErrInternalServer)
		return
	}

	render.JSON(w, r, users)
}

func (a *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	if userID == "" {
		render.Render(w, r, ErrBadParams)
		return
	}

	user, err := a.userService.GetByID(userID)
	if err != nil {
		render.Render(w, r, ErrInternalServer)
		return
	}

	render.JSON(w, r, user)
}
