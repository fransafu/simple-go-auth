package handlers

import (
	"github.com/fransafu/simple-go-auth/internal/app/services"
)

type Handlers struct {
	AuthHandler IAuthHandler
	UserHandler IUserHandler
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		AuthHandler: NewAuthHandler(),
		UserHandler: NewUserHandler(services.UserService),
	}
}
