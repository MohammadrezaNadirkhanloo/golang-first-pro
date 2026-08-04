package bootstrap

import (
	"first-app/internal/modules/home/routes"
	"first-app/pkg/config"
	"first-app/pkg/database"
	"first-app/pkg/html"
	"first-app/pkg/routing"
	"first-app/pkg/sessions"
	"first-app/pkg/static"
	"fmt"
	"log"
)

func Server() {
	config.Set()

	database.Connect()

	configs := config.Get()

	routing.Init()
	
	router := routing.GetRouter()
	sessions.Start(router)

	static.LoadStatic(router)

	html.LoadHTML(router)

	routes.Routes(router)

	if err := router.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
