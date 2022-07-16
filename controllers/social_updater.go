package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UpdaterSocial(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contenType := helpers.GetContentType(c)
	Social := models.Social{}

	socialID, _ := strconv.Atoi(c.Param("socialID"))
	id := uint(userData["id"].(float64))

	if contenType == appJSON {
		c.ShouldBindJSON(&Social)
	} else {
		c.ShouldBind(&Social)
	}

	Social.ID = id

	err := db.Model(&Social).Where("id = ?", socialID).Updates(models.Social{
		Name:           Social.Name,
		SocialMediaUrl: Social.SocialMediaUrl,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Social)
}
