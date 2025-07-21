package client

import (
	authpb "github.com/Babushkin05/simple-marketplace/api-gateway/api/gen/auth"
	"google.golang.org/grpc"
)

type AuthClient interface {
	authpb.AuthServiceClient
}

func NewAuthClient(conn *grpc.ClientConn) AuthClient {
	return authpb.NewAuthServiceClient(conn)
}
