package main

import (
	"log"

	"github.com/ivanpaghubasan/apcp-hub-backend/internal/config"
	"github.com/ivanpaghubasan/apcp-hub-backend/internal/database"
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

    
}
