package parser

import (
	"os"
	"testing"

	"github.com/bebrochkas/rural_potatoes/core/config"
	"github.com/bebrochkas/rural_potatoes/core/internal/db"
)

func TestMain(m *testing.M) {
	config.Initialize("../../../.env")
	db.Initialize()
	os.Exit(m.Run())
}

func TestFetchFilm(t *testing.T) {
	type args struct {
		filmId int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test movie 300", args{325}, false},
		{"Test movie 301", args{303}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := fetchFilm(tt.args.filmId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchFilm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
