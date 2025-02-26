package parser

import (
	"github.com/bebrochkas/rural_potatoes/core/internal/db/films"
	"github.com/charmbracelet/log"
)

func FetchBatch(size int) (int, error) {
	var fetched int
	for filmId := 1340709; filmId <= (1340709 + size); filmId++ {
		kinoFilm, err := fetchFilm(filmId)
		if err != nil {
			log.Error("filed to fetch film with", "id", filmId, "and err", err)
		}

		if kinoFilm.Name == "" || kinoFilm.Desc == "" {
			log.Error("empty name or description for film with", "id", filmId)
			continue
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
