package tags

import (
	"testing"

	"github.com/bebrochkas/rural_potatoes/core/config"
	"github.com/bebrochkas/rural_potatoes/core/internal/db"
	"os"
)

func TestMain(m *testing.M) {
	config.Initialize("../../../../.env")
	db.Initialize()
	os.Exit(m.Run())
}
