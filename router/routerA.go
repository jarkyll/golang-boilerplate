package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRouterA(router *gin.Engine) *gin.Engine {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())


	router.GET("pong", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	return router
}