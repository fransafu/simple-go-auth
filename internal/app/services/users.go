package services

import (
	"github.com/fransafu/simple-go-auth/internal/app/models"
)

type IUserService interface {
	GetAll() (*[]models.User, error)
	GetByID(userID string) (*models.User, error)
}

type UserService struct {
	userRepository models.UserRepository
}

func NewUserService(userRepository models.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) GetAll() (*[]models.User, error) {
	users, err := u.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *UserService) GetByID(userID string) (*models.User, error) {
	user, err := u.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
