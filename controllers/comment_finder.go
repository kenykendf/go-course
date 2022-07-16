package controllers

import (
	"final-project/database"
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FinderComment(c *gin.Context) {
	var comment []models.Comment

	db := database.GetDB()

	err := db.Preload("User").Preload("Photo").Find(&comment).Error
	// err := database.DB.Find(&comment).Error
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    comment,
	})
}
