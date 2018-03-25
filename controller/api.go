package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konojunya/spotify-api-sample/service"
)

func GetMyTracks(c *gin.Context) {
	tracks, err := service.GetMyTracks()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, tracks)
}
