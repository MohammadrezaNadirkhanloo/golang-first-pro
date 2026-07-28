package bootstrap

import (
	"first-app/internal/database/midration"
	"first-app/pkg/config"
	"first-app/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()
	midration.Migrate()
}
