package db

import "github.com/bebrochkas/rural_potatoes/core/models"

var systemTags = []string{}

func MigrateTags() (*[]*models.Tag, error) {

	var tags []*models.Tag

	for _, systemTag := range systemTags {
		tags = append(tags, &models.Tag{Name: systemTag})
	}

	// thematic tags loading logic

	return &tags, nil

}
