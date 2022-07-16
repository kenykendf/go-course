package controllers

import (
	"final-project/database"
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FinderSocial(c *gin.Context) {
	var social []models.Social

	err := database.DB.Find(&social).Error
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    social,
	})
}
