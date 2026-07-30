package bootstrap

import (
	"first-app/internal/database/seeder"
	"first-app/pkg/config"
	"first-app/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()
	seeder.Seed()
}
