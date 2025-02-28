package db

import (
	"bufio"
	"os"

	"encoding/csv"
	"strconv"

	"github.com/bebrochkas/rural_potatoes/core/config"
	"github.com/bebrochkas/rural_potatoes/core/models"
)

func readCSV(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	return records
}

func toUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}

func autoFilms() error {
	records := readCSV(config.Cfg.MIGRATION_PREFIX + "films.csv")

	for i, record := range records {
		if i == 0 {
			continue
		}

		film := models.Film{
			ID:           toUint(record[0]),
			Title:        record[1],
			Description:  record[2],
			PosterPreUrl: record[3],
			PosterUrl:    record[4],
			BackdropUrl:  record[5],
		}
		DB.Save(&film)
	}

	return nil
}

func autoTags() error {
	records := readCSV(config.Cfg.MIGRATION_PREFIX + "tags.csv")

	for i, record := range records {
		if i == 0 {
			continue
		}

		tag := models.Tag{
			ID:   toUint(record[0]),
			Name: record[1],
			Hex:  record[2],
			Type: record[3],
		}
		DB.Save(&tag)
	}

	return nil
}

func autoFilmsTags() error {
	records := readCSV(config.Cfg.MIGRATION_PREFIX + "filmtags.csv")
	for _, record := range records[1:] { // Skip header
		tagID := toUint(record[0])
		filmID := toUint(record[1])

		DB.Exec("INSERT INTO film_tags (tag_id, film_id) VALUES (?, ?) ON CONFLICT DO NOTHING", tagID, filmID)
	}

	return nil
}

func ColoriseTags(prefix string) error {

	// thematic

	tagDir, err := os.ReadDir(prefix)
	if err != nil {
		return err
	}

	for _, tagFilePath := range tagDir {
		tagFile, err := os.Open(prefix + tagFilePath.Name())
		if err != nil {
			return err
		}
		defer tagFile.Close()

		var hex string
		scanner := bufio.NewScanner(tagFile)
		if scanner.Scan() {
			hex = scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			return err
		}

		name := tagFilePath.Name()[:len(tagFilePath.Name())-4]

		res := DB.Where("name=?", name).FirstOrCreate(&models.Tag{Name: name, Hex: hex, Type: "thematic"}, &models.Tag{Name: name})

		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected > 0 {
			continue
		}

		if err := DB.Model(&models.Tag{}).Where("name=?", name).Updates(models.Tag{Hex: hex, Type: "thematic"}).Error; err != nil {
			return err
		}
	}

	return nil
}
