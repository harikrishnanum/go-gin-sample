package users

import (
	"go-gin-sample/api/users/services"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := services.GetUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(200, users)
}
