package grpc

import (
	"fmt"
	"net"

	pb "github.com/Babushkin05/simple-marketplace/auth-service/api/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer(handler pb.AuthServiceServer, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, handler)
	reflection.Register(s)

	fmt.Printf("gRPC server started on port %d\n", port)
	return s.Serve(lis)
}
