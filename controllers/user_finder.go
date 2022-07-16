package controllers

import (
	"final-project/database"
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FinderUser(c *gin.Context) {
	var user []models.User

	err := database.DB.Find(&user).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "finder pong",
		"result":  user,
	})
}
