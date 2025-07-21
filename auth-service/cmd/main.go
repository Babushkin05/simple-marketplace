package main

import (
	"log"

	cfgPkg "github.com/Babushkin05/simple-marketplace/auth-service/internal/config"
	"github.com/Babushkin05/simple-marketplace/auth-service/internal/db"
	grpcServer "github.com/Babushkin05/simple-marketplace/auth-service/internal/grpc"
	"github.com/Babushkin05/simple-marketplace/auth-service/internal/service"
)

func main() {
	// Load config
	cfg := cfgPkg.MustLoad()
	log.Println("Config loaded")

	// Connect to database
	dbConn := db.MustInitPostgres(*cfg)
	defer dbConn.Close()

	// Init dependencies
	repo := db.NewPostgresRepo(dbConn)
	authService := service.NewAuthService(repo, cfg.JWT.Secret, cfg.JWT.TTL)
	handler := grpcServer.NewAuthHandler(authService)

	// Start gRPC server
	if err := grpcServer.StartGRPCServer(handler, cfg.Server.GRPCPort, cfg.Server.Host); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
}
