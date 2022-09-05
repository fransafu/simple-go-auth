package models

import (
	"errors"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByUsernameAndPassword(username string, password string) (*User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(userID string) (User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

type User struct {
	gorm.Model
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty" gorm:"unique"`
	Password string `json:"password,omitempty"`
}

func (User) TableName() string {
	return "users"
}

func (c *UserRepository) GetUserByUsernameAndPassword(username string, password string) (User, error) {
	var user User

	err := c.DB.First(&user, "username = ? AND password = ?", username, password).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (c *UserRepository) GetAllUsers() ([]User, error) {
	var user []User

	err := c.DB.Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (c *UserRepository) GetUserByID(userID string) (*User, error) {
	var user User

	err := c.DB.First(&user, "id = ?", userID).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &user, nil
}
