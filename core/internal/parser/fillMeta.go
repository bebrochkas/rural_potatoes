package parser

import (
	"fmt"

	"github.com/bebrochkas/rural_potatoes/core/internal/pb"
	"github.com/bebrochkas/rural_potatoes/core/models"
)

func proccesMeta(kinoFilm *KinoPoiskResp) (models.Film, []models.Tag, error) {

	film := models.Film{
		Title:        kinoFilm.Name,
		Description:  kinoFilm.Desc,
		PosterPreUrl: kinoFilm.Poster.PreviewUrl,
		PosterUrl:    kinoFilm.Poster.Url,
		BackdropUrl:  kinoFilm.Backdrop.Url}

	var tags []models.Tag

	// releaseYear
	var realeseTag string

	switch {

	case kinoFilm.Year <= 1959:
		realeseTag = "1930-1959"

	case kinoFilm.Year <= 1979:
		realeseTag = "1960-1979"

	case kinoFilm.Year <= 1999:
		realeseTag = "1980-1999"

	case kinoFilm.Year <= 2009:
		realeseTag = "2000-2009"

	case kinoFilm.Year <= 2019:
		realeseTag = "2010-2019"

	default:
		realeseTag = "позже 2019"

	}
	tags = append(tags, models.Tag{Name: realeseTag, Type: "release"})

	// countries
	for _, country := range kinoFilm.Countries {
		tags = append(tags, models.Tag{Name: country.Name, Type: "country"})
	}

	// age rating
	tags = append(tags, models.Tag{Name: fmt.Sprintf("%d+", kinoFilm.AgeRate), Type: "age rating"})

	// critics rate
	var rateTag string

	var susp_cnt float32
	var totalRating float32
	var susp_rating float32

	if kinoFilm.Rating.KP != 0 {
		susp_cnt++
		totalRating += kinoFilm.Rating.KP
	}

	if kinoFilm.Rating.IMDB != 0 {
		susp_cnt++
		totalRating += kinoFilm.Rating.IMDB
	}

	if susp_cnt > 0 {
		susp_rating = totalRating / susp_cnt
		switch {
		case susp_rating <= 3:
			rateTag = "Низкие оценки"
		case susp_rating <= 6:
			rateTag = "Средние оценки"
		case susp_rating <= 8.5:
			rateTag = "Выскоие оценки"
		default:
			rateTag = "Крайне выскоие оценки"
		}
	}

	film.Rate = susp_rating

	theme_tags, err := pb.GetTags(film.Description)

	if err != nil {
		return film, tags, err
	}

	for _, theme_tag := range theme_tags {
		tags = append(tags, models.Tag{Name: theme_tag, Type: "thematic"})
	}

	tags = append(tags, models.Tag{Name: rateTag, Type: "rate"})

	return film, tags, nil

}
