package router

import (
	"github.com/gin-gonic/gin"
	"github.com/konojunya/spotify-api-sample/controller"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	// static files
	r.Static("/js", "./public/js")
	r.Static("/css", "./public/css")
	r.Static("/image", "./public/image")

	r.LoadHTMLGlob("view/*")

	r.GET("/", controller.Index)
	r.GET("/login", controller.Login)
	r.GET("/oauth", controller.CallBack)

	api := r.Group("/api")
	apiRouter(api)

	return r
}
