package controllers

import (
	"assignment-2/database"
	"assignment-2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	var newOrder models.Order

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := database.DB.Create(&newOrder).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func GetOrder(ctx *gin.Context) {
	var result []models.Order

	err := database.DB.Preload("Items").Find(&result).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func UpdateOrder(ctx *gin.Context) {
	var updateOrder models.Order

	var orderID string

	orderID = ctx.Param("orderID")
	id, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := ctx.ShouldBindJSON(&updateOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	updateOrder.ID = uint(id)

	err = database.DB.Save(&updateOrder).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for i := 0; i < len(updateOrder.Items); i++ {
		err = database.DB.Save(&updateOrder.Items[i]).Error
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	ctx.JSON(http.StatusOK, updateOrder)
}

func DeleteOrder(ctx *gin.Context) {
	var deleteOrder models.Order

	orderID := ctx.Param("orderID")
	id, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	deleteOrder.ID = uint(id)

	err = database.DB.Delete(&deleteOrder).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
