package main

import (
	"communications/internal/config"
	"communications/internal/handlers"
	"communications/internal/server"
)

func main() {
	cfg := config.Load()
	router := handlers.Init(cfg)
	server.Listen(cfg.Port, router)
}
