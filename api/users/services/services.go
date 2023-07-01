package services

import (
	"go-gin-sample/api/users/models"
	"go-gin-sample/logger"
)

func GetUsers() ([]models.User, error) {
	logger.Log.Info("Fetching users from the database")
	// Simulated implementation fetching users from the database
	users := []models.User{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Jane"},
	}

	return users, nil
}
