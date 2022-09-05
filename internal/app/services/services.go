package services

import (
	"github.com/fransafu/simple-go-auth/internal/app/database"
	"github.com/fransafu/simple-go-auth/internal/app/models"
)

type Services struct {
	UserService IUserService
}

func NewServices() *Services {
	db, _ := database.GetConnection()

	return &Services{
		UserService: NewUserService(models.UserRepository{DB: db}),
	}
}
