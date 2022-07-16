package controllers

import (
	"final-project/database"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RemoverPhoto(c *gin.Context) {
	var removePhoto models.Photo

	photoID := c.Param("paramID")
	id, err := strconv.ParseUint(photoID, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	removePhoto.ID = uint(id)

	err = database.DB.Delete(&removePhoto).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo's Account Has Been Deleted",
	})
}
