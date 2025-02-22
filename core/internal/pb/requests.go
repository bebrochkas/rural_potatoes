package pb

import (
	"context"

	"github.com/bebrochkas/rural_potatoes/core/proto"
)

func GetTags(description string) ([]string, error) {

	resp, err := client.Tag(context.Background(), &proto.TagRequest{Description: description})

	if err != nil {
		return nil, err
	}

	return resp.Tags, nil

}
