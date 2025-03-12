package main

import (
	"log"
	"net/http"
	routes "song-library/internal/api/routes"
	clients "song-library/internal/clients"
	repositories "song-library/internal/repositories"
	services "song-library/internal/services"
	utils "song-library/internal/utils"
)

func main() {
	// loading configuration
	cfg, err := utils.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	// creating logger
	log := utils.NewLogger("Song library: ", cfg.Logger.EnableDebug)

	// creation repository
	repo, err := repositories.NewRepository(cfg, log)
	if err != nil {
		log.Fatal("Error creating repository: ", err)
	}
	log.Info("Repository created successful")

	// start migrations
	if err := repo.RunMigrations(cfg); err != nil {
		log.Warn("Error running migrations: ", err)
	}
	log.Info("Migrations applied successfully")

	// client creation
	client, err := clients.NewExternalAPIClient(cfg, log)
	if err != nil {
		log.Fatal("Error creating client: ", err)
	}
	log.Info("Client created successful")

	// service creation
	service, err := services.NewService(client, repo, log)
	if err != nil {
		log.Fatal("Error creating service: ", err)
	}
	log.Info("Service created successful")

	// creating routes
	router, err := routes.CreateRoutes(service)
	if err != nil {
		log.Fatal("Error creating routes: ", err)
	}
	log.Info("Routes created successful")

	// starting http server
	log.Info("Starting the server on port ", cfg.Server.Port)
	if err := http.ListenAndServe(":"+cfg.Server.Port, router); err != nil {
		log.Fatal("Error starting server: ", err)
	}
	log.Info("Server started successful")
}
