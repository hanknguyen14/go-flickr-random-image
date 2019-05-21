package flickr

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type (
	photo struct {
		ID             string `json:"id"`
		Owner          string `json:"owner"`
		Secret         string `json:"secret"`
		Server         string `json:"server"`
		Farm           int    `json:"farm"`
		Title          string `json:"title"`
		IsPublic       int    `json:"ispublic"`
		Isfriend       int    `json:"isfriend"`
		Isfamily       int    `json:"isfamily"`
		Originalsecret string `json:"originalsecret"`
		Originalformat string `json:"originalformat"`
		URL            string `json:"url_l"`
		Height         string `json:"height_l"`
		Width          string `json:"width_l"`
	}
	photos struct {
		Page    int     `json:"page"`
		Pages   int     `json:"pages"`
		Perpage int     `json:"perpage"`
		Total   string  `json:"total"`
		Photo   []photo `json:"photo"`
	}

	photoRaw struct {
		Photos photos `json:"photos"`
		Stat   string `json:"stat"`
	}

	randomPhoto struct {
		Title    string
		URL      string
		Filetype string
	}
)

func (Client *client) getRandomImage(UserID string, PerPage int) (*randomPhoto, error) {
	response, errors := Client.request(params{
		"user_id":  UserID,
		"media":    "photos",
		"extras":   "url_l,original_format",
		"per_page": fmt.Sprintf("%d", PerPage),
	})

	if errors != nil {
		return nil, errors
	}

	photoRaw := &photoRaw{}

	jsonErr := json.Unmarshal(response, photoRaw)

	if jsonErr != nil {
		return nil, jsonErr
	}

	randomNumber := rand.Intn(50)

	return &randomPhoto{
		Title:    photoRaw.Photos.Photo[randomNumber].Title,
		URL:      photoRaw.Photos.Photo[randomNumber].URL,
		Filetype: photoRaw.Photos.Photo[randomNumber].Originalformat,
	}, nil
}
