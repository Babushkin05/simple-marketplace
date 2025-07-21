package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Babushkin05/simple-marketplace/api-gateway/internal/client"
	"github.com/Babushkin05/simple-marketplace/api-gateway/internal/config"
	"github.com/Babushkin05/simple-marketplace/api-gateway/internal/handler"
)

func main() {
	cfg := config.MustLoad()
	log.Println("Config loaded")

	// Init gRPC connections
	authConn, goodsConn, err := client.InitConnections(*cfg)
	if err != nil {
		log.Fatalf("failed to initialize gRPC connections: %v", err)
	}
	defer authConn.Close()
	defer goodsConn.Close()

	// Init Gin handler
	router := handler.NewHandler(authConn, goodsConn)

	addr := ":" + strconv.Itoa(cfg.Server.Port)
	log.Printf("Starting HTTP server on %s", addr)
	if err := router.Run(addr); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
