package grpc

import (
	"context"

	pb "github.com/Babushkin05/simple-marketplace/auth-service/api/gen"
	"github.com/Babushkin05/simple-marketplace/auth-service/internal/service"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	svc *service.AuthService
}

func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user, login, err := h.svc.Register(ctx, req.Login, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		UserId: user,
		Login:  login,
	}, nil
}

func (h *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := h.svc.Login(ctx, req.Login, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{Token: token}, nil
}

func (h *AuthHandler) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	user, err := h.svc.ValidateToken(ctx, req.Token)
	if err != nil {
		return &pb.ValidateTokenResponse{Valid: false}, nil
	}

	return &pb.ValidateTokenResponse{
		UserId: user.ID,
		Login:  user.Login,
		Valid:  true,
	}, nil
}
