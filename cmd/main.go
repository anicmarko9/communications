package main

import (
	"communications/internal/config"
	"communications/internal/database"
	"communications/internal/handlers"
	"communications/internal/server"
)

// Entry point for the application.
func main() {
	cfg := config.Load()
	db := database.Connect(cfg)
	router := handlers.Init(cfg, db)
	server.Listen(cfg.Port, router)
}
