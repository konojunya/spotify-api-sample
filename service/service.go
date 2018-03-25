package service

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/konojunya/spotify-api-sample/auth"
	"github.com/makki0205/gojwt"
	"golang.org/x/oauth2"
)

var (
	c *http.Client
)

type Client struct {
	client *http.Client
}

func ClientMiddleware(ctx *gin.Context) {
	t, err := ctx.Cookie("spotify-sample-jwt")
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
	if c != nil {
	}
	payload, err := jwt.Decode(t)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	token := oauth2.Token{
		AccessToken:  payload["accessToken"],
		RefreshToken: payload["refreshToken"],
	}

	c = auth.GetClient(&token)
	ctx.Next()
}

func getClient() Client {
	return Client{
		client: c,
	}
}

func (c Client) Get(url string) ([]byte, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
