package parser

import (
	"github.com/charmbracelet/log"
	"github.com/bebrochkas/rural_potatoes/core/internal/db/films"
)

func FetchBatch(size int) (int, error) {
	var fetched int
	for filmId := 298; filmId <= (298 + size); filmId++ {
		kinoFilm, err := fetchFilm(filmId)
		if err != nil {
			log.Error("filed to fetch film with", "id", filmId, "and err", err)
		}

		if kinoFilm.Name == "" || kinoFilm.Desc == "" {
			return 0, nil
		}

		film, tags, err := proccesMeta(kinoFilm)

		if err != nil {
			log.Error("filed to process film with", "id", filmId, "and err", err)
			continue
		}

		err = films.InsertFilmWTags(&film, tags)

		if err != nil {
			log.Error("filed to save film to db with", "id", filmId, "and err", err)
			continue
		}

		log.Info("succefuly created film", "with tittle", film.Title)

	}
	return fetched, nil
}
