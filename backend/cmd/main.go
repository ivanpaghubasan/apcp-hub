package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ivanpaghubasan/apcp-hub-backend/internal/config"
	"github.com/ivanpaghubasan/apcp-hub-backend/internal/database"
	"github.com/ivanpaghubasan/apcp-hub-backend/internal/repository"
	"github.com/ivanpaghubasan/apcp-hub-backend/internal/routes"
	"github.com/ivanpaghubasan/apcp-hub-backend/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	repos := repository.NewRepositories(db)

	services := service.NewServices(repos)

	r := gin.Default()

	routes.SetupRouter(r, services)

	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
