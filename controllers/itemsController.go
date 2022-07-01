package controllers

import (
	"assignment-2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateItem(ctx *gin.Context) {
	var newItem models.Item

	if err := ctx.ShouldBindJSON(&newItem); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

}

func GetItem(ctx *gin.Context) {
	//
}

func UpdateItem(ctx *gin.Context) {
	//
}

func DeleteItem(ctx *gin.Context) {
	//
}
