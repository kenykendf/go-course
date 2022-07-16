package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreatorSocial(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Social := models.Social{}
	userData := c.MustGet("userData").(jwt.MapClaims)

	id := uint(userData["id"].(float64))

	Social.UserID = id

	if contentType == appJSON {
		c.ShouldBindJSON(&Social)
	} else {
		c.ShouldBind(&Social)
	}

	err := db.Debug().Create(&Social).Error
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": Social,
	})
}
