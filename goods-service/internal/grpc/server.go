package grpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	authpb "github.com/Babushkin05/simple-marketplace/goods-service/api/gen/auth"
	goodspb "github.com/Babushkin05/simple-marketplace/goods-service/api/gen/goods"
	"github.com/Babushkin05/simple-marketplace/goods-service/internal/config"
	"github.com/Babushkin05/simple-marketplace/goods-service/internal/service"
)

func RunGRPCServer(cfg config.Config, srv service.GoodsService, authClient authpb.AuthServiceClient) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	grpcServer := grpc.NewServer()

	handler := NewGoodsHandler(srv, authClient)
	goodspb.RegisterGoodsServiceServer(grpcServer, handler)

	log.Printf("gRPC server listening at %s:%d", cfg.Server.Host, cfg.Server.Port)
	return grpcServer.Serve(listener)
}
