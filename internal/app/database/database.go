package database

import (
	"github.com/fransafu/simple-go-auth/internal/app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Initialize() error {
	db, err := GetConnection()
	if err != nil {
		return err
	}

	db.AutoMigrate(&models.User{})

	return nil
}

func GetConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
