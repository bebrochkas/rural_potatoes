package main

import (
	"flag"
	"path/filepath"

	"github.com/bebrochkas/rural_potatoes/core/config"
	"github.com/bebrochkas/rural_potatoes/core/internal/api"
	"github.com/bebrochkas/rural_potatoes/core/internal/db"
	// "github.com/bebrochkas/rural_potatoes/core/internal/parser"
	"github.com/bebrochkas/rural_potatoes/core/internal/pb"
	"github.com/charmbracelet/log"
	"os"
)

func main() {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	log.Print(exPath)

	envPath := flag.String("env", "../.env", "provide .env filepath")
	flag.Parse()

	config.Initialize(*envPath)

	if err := db.Initialize(true); err != nil {
		log.Fatal("failed to init DB with", "err", err)
	}

	if err := pb.Initialize(); err != nil {
		log.Fatal("failed to init PB with", "err", err)
	}

	// parser.FetchBatch(24)

	api.Initialize()

}
