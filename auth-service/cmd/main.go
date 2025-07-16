package main

import (
	"log"

	"github.com/Babushkin05/simple-marketplace/auth-service/internal/config"
)

func main() {
	cfg := config.MustLoad()
	log.Printf("Loaded config: %s", cfg)

	// init db + migration

	// init grpc server

}
