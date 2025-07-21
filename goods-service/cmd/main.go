package main

import (
	"log"

	"github.com/Babushkin05/simple-marketplace/goods-service/internal/config"
	"github.com/Babushkin05/simple-marketplace/goods-service/internal/db"
	"github.com/Babushkin05/simple-marketplace/goods-service/internal/grpc"
	"github.com/Babushkin05/simple-marketplace/goods-service/internal/service"
)

func main() {
	// Load config
	cfg := config.MustLoad()
	log.Printf("Config loaded")

	// Init DB
	dbConn := db.MustInitPostgres(*cfg)
	defer dbConn.Close()
	log.Println("Connected to DB")

	// Init auth client
	authClient, err := grpc.NewAuthClient(*cfg)
	if err != nil {
		log.Fatalf("Failed to connect to auth service: %v", err)
	}

	// Init service layer
	repo := db.NewPostgresRepo(dbConn)
	svc := service.NewService(repo, authClient)

	// Start gRPC server
	if err := grpc.RunGRPCServer(*cfg, svc, authClient); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	log.Printf("gRPC server started on %s:%d", cfg.Server.Host, cfg.Server.Port)
}
