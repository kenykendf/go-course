package router

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.POST("/user/login", controllers.LoginUser)
	r.GET("/user", controllers.FinderUser)
	r.POST("/user", controllers.CreatorUser)

	r.Use(middlewares.Authentication())
	r.PUT("/user/:paramID", middlewares.UserAuthorization(), controllers.UpdaterUser)
	r.DELETE("/user/:paramID", middlewares.UserAuthorization(), controllers.RemoverUser)

	r.GET("/photo", controllers.FinderPhoto)
	r.POST("/photo", controllers.CreatorPhoto)
	r.PUT("/photo/:paramID", middlewares.UserAuthorization(), controllers.UpdaterPhoto)
	r.DELETE("/photo/:paramID", middlewares.UserAuthorization(), controllers.RemoverPhoto)

	r.GET("/comment", controllers.FinderComment)
	r.POST("/comment", controllers.CreatorComment)
	r.PUT("/comment/:paramID", controllers.UpdaterComment)
	r.DELETE("/comment/:paramID", controllers.RemoverComment)

	r.GET("/social", controllers.FinderSocial)
	r.POST("/social", controllers.CreatorSocial)
	r.PUT("/social/:paramID", controllers.UpdaterSocial)
	r.DELETE("/social/:paramID", controllers.RemoverSocial)

	return r
}
