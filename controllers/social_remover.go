package controllers

import (
	"final-project/database"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RemoverSocial(c *gin.Context) {
	var removeSocial models.Social

	socialID := c.Param("socialID")
	id, err := strconv.ParseUint(socialID, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	removeSocial.ID = uint(id)

	err = database.DB.Delete(&removeSocial).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data Has Been Deleted",
	})
}
