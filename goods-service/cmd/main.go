package main

import (
	"log"

	"github.com/Babushkin05/simple-marketplace/goods-service/internal/config"
)

func main() {
	cfg := config.MustLoad()
	log.Printf("Config has loaded")

	// init db + migration

	// init grpc server

}
