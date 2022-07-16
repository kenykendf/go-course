package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreatorComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Comment := models.Comment{}
	userData := c.MustGet("userData").(jwt.MapClaims)

	id := uint(userData["id"].(float64))

	Comment.UserID = id
	Comment.PhotoID = 1

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Debug().Create(&Comment).Error
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Comment Created",
		"data":    Comment,
	})
}
