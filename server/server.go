package server

import (
	"go-gin-sample/api/users"
	"go-gin-sample/logger"

	"github.com/gin-gonic/gin"
)

func Run() error {
	router := gin.Default()
	router.GET("/users", users.GetUsers)

	logger.Log.Info("Starting the server on port 9000")
	err := router.Run(":9000")
	if err != nil {
		logger.Log.Errorf("Failed to start the server: %v", err)
		return err
	}
	return nil
}
