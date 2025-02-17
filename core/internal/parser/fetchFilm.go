package parser

import (
	"encoding/json"
	"fmt"
	"github.com/bebrochkas/rural_potatoes/core/config"
	"io"
	"net/http"
)

var client = &http.Client{}

// "poster": {
//    "url": "https://image.openmoviedb.com/kinopoisk-images/10671298/72d3ef84-c785-46aa-8a8d-16e0f0847c56/orig",
//    "previewUrl": "https://image.openmoviedb.com/kinopoisk-images/10671298/72d3ef84-c785-46aa-8a8d-16e0f0847c56/x1000"
//  },
//  "backdrop": {
//    "url": "https://image.openmoviedb.com/kinopoisk-ott-images/224348/2a0000018f6e9e449af3e7e1c0e564125807/orig",
//    "previewUrl": "https://image.openmoviedb.com/kinopoisk-ott-images/224348/2a0000018f6e9e449af3e7e1c0e564125807/x1000"
//  },

type KinoPoiskResp struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Year      uint   `json:"year"`
	AgeRate   uint   `json:"ageRating"`
	Desc      string `json:"description"`
	ShortDesc string `json:"shortDescription"`
	Countries []struct {
		Name string `json:"name"`
	} `json:"countries"`
	Poster []struct {
		Url        string `json:"url"`
		PreviewUrl string `json:"previewUrl"`
	} `json:"poster"`
	Backdrop []struct {
		Url string `json:"url"`
	} `json:"backdrop"`
	Rating struct {
		KP   float32 `json:"kp"`
		IMDB float32 `json:"imdb"`
		FC   float32 `json:"filmCritics"`
	} `json:"rating"`
}

func fetchFilm(filmId int) (*KinoPoiskResp, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.kinopoisk.dev/v1.4/movie/%d", filmId), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-API-KEY", config.Cfg.KINOPOISK_KEY)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var movie KinoPoiskResp
	if err := json.Unmarshal(body, &movie); err != nil {
		return nil, err
	}

	return &movie, nil
}
