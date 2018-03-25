package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konojunya/spotify-api-sample/auth"
)

func Index(c *gin.Context) {
	jwt, err := c.Cookie("spotify-sample-jwt")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if jwt != "" {
		c.HTML(http.StatusOK, "main.html", nil)
		return
	}

	c.HTML(http.StatusOK, "index.html", nil)
}

func Login(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, auth.GetRedirectURL())
}

func CallBack(c *gin.Context) {
	code := c.Query("code")
	jwt, err := auth.GetJWTToken(code)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.SetCookie("spotify-sample-jwt", jwt, 1000*60*60*24*7, "/", "http://localhost:9000", false, false)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
