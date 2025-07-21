package grpc

import (
	"fmt"

	"google.golang.org/grpc"

	authpb "github.com/Babushkin05/simple-marketplace/goods-service/api/gen/auth"
	"github.com/Babushkin05/simple-marketplace/goods-service/internal/config"
)

func NewAuthClient(cfg config.Config) (authpb.AuthServiceClient, error) {
	conn, err := grpc.Dial(cfg.AuthService.Address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to auth-service: %w", err)
	}
	return authpb.NewAuthServiceClient(conn), nil
}
