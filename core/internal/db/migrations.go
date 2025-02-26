package db

import (
	"bufio"
	"github.com/bebrochkas/rural_potatoes/core/models"
	"os"
)

func MigrateTags() error {
	var tags []*models.Tag

	prefix := "../tagger/tags_data/"

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

		// Prepare the tag data
		tag := &models.Tag{Name: tagFilePath.Name()[:len(tagFilePath.Name())-4], Hex: hex}

		if err := DB.FirstOrCreate(&tag, models.Tag{Name: tag.Name}).Error; err != nil {
			return err
		}

		tags = append(tags, tag)
	}

	return nil
}
