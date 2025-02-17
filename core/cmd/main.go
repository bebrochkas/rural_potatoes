package main

import (
	"flag"

	"github.com/bebrochkas/rural_potatoes/core/config"
	"github.com/bebrochkas/rural_potatoes/core/internal/api"
	"github.com/bebrochkas/rural_potatoes/core/internal/db"
	"github.com/bebrochkas/rural_potatoes/core/internal/parser"
	"github.com/charmbracelet/log"
)

func main() {

	envPath := flag.String("env", "../.env", "provide .env filepath")
	flag.Parse()

	config.Initialize(*envPath)

	if err := db.Initialize(); err != nil {
		log.Fatal("failed to init DB with", "err", err)
	}

	api.Initialize()

	parser.FetchBatch(0)

}
