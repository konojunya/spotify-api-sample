package main

import "github.com/konojunya/spotify-api-sample/router"

func main() {
	r := router.GetRouter()
	r.Run(":9000")
}
