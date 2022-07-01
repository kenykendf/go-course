package routers

import (
	"assignment-2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/order", controllers.CreateOrder)

	router.PUT("/order/:orderID", controllers.UpdateOrder)

	router.GET("/orders", controllers.GetOrder)

	router.DELETE("/order/:orderID", controllers.DeleteOrder)

	// router.POST("/item", controllers.CreateItem)

	// router.PUT("/item/:itemID", controllers.UpdateItem)

	// router.GET("/items/:itemID", controllers.GetItem)

	// router.DELETE("/item/:itemID", controllers.DeleteItem)

	return router
}
