package controllers

import (
	"final-project/database"
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FinderPhoto(c *gin.Context) {
	var photo []models.Photo

	db := database.GetDB()

	err := db.Preload("User").Find(&photo).Error
	// err := database.DB.Find(&photo).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    photo,
	})
}
