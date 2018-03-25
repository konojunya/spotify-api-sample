package service

import (
	"encoding/json"
	"log"

	"github.com/konojunya/spotify-api-sample/model"
)

func GetMyTracks() (*model.Response, error) {
	client := getClient()
	body, err := client.Get("https://api.spotify.com/v1/me/tracks")
	if err != nil {
		return nil, err
	}

	var response model.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &response, nil
}
