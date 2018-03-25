package router

import (
	"github.com/gin-gonic/gin"
	"github.com/konojunya/spotify-api-sample/controller"
	"github.com/konojunya/spotify-api-sample/service"
)

func apiRouter(api *gin.RouterGroup) {
	api.Use(service.ClientMiddleware)
	api.GET("/tracks", controller.GetMyTracks)
}
