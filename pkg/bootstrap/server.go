package bootstrap

import (
	"first-app/internal/modules/home/routes"
	"first-app/pkg/config"
	"first-app/pkg/routing"
	"fmt"
	"log"
)

func Server() {
	config.Set()

	configs := config.Get()

	routing.Init()

	router := routing.GetRouter()

	routes.Routes(router)

	if err := router.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
