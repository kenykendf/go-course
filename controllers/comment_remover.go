package controllers

import (
	"final-project/database"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RemoverComment(c *gin.Context) {
	var removeComment models.Comment

	commentID := c.Param("commentID")
	id, err := strconv.ParseUint(commentID, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	removeComment.ID = uint(id)

	err = database.DB.Delete(&removeComment).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment Has Been Deleted",
	})
}
