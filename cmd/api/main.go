package main

import (
	"context"
	"fmt"
	"log"

	"lms-core-api/internal/config"
	"lms-core-api/internal/database"
	"lms-core-api/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load configuration: %v", err)
	}

	_, sqlDB, err := database.Open(context.Background(), cfg.Database)
	if err != nil {
		log.Fatalf("connect to database: %v", err)
	}
	defer sqlDB.Close()

	router := server.New(sqlDB)

	if err := router.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		log.Fatalf("run server: %v", err)
	}
}
