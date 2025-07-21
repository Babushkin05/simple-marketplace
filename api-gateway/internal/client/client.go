package client

import (
	"fmt"

	"github.com/Babushkin05/simple-marketplace/api-gateway/internal/config"
	"google.golang.org/grpc"
)

func InitConnections(cfg config.Config) (*grpc.ClientConn, *grpc.ClientConn, error) {
	authConn, err := grpc.Dial(cfg.AuthService.Addr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to auth service: %w", err)
	}

	goodsConn, err := grpc.Dial(cfg.GoodsService.Addr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to goods service: %w", err)
	}

	return authConn, goodsConn, nil
}
