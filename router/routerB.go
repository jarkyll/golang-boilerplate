package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRouterB(router *gin.Engine) *gin.Engine {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())


	router.GET("pang", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "pang"})
	})

	return router
}