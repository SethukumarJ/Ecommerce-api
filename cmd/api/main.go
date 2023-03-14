package main

import (
	"log"

	config "ecommerce/pkg/config"
	di "ecommerce/pkg/di"
)


// @title Go + Gin Ecommerce API
// @version 1.0
// @description This is a sample server Job Portal server. You can visit the GitHub repository at https://github.com/fazilnbr/Job_Portal_Project

// @contact.name API Support
// @contact.url https://fazilnbr.github.io/mypeosolal.web.portfolio/
// @contact.email fazilkp2000@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:3000
// @BasePath /
// @query.collection.format multi
func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
