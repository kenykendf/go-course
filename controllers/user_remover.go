package controllers

import (
	"final-project/database"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RemoverUser(c *gin.Context) {
	var removeUser models.User

	userID := c.Param("userID")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	removeUser.ID = uint(id)

	err = database.DB.Delete(&removeUser).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User's Account Has Been Deleted",
	})
}
