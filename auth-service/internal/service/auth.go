package service

import (
	"context"
	"errors"
	"time"

	"github.com/Babushkin05/simple-marketplace/auth-service/internal/auth"
	"github.com/Babushkin05/simple-marketplace/auth-service/internal/db"
	"github.com/google/uuid"
)

type AuthService struct {
	Repo      db.UserRepository
	JWTSecret string
	TokenTTL  time.Duration
}

func NewAuthService(repo db.UserRepository, secret string, tokenTTL time.Duration) *AuthService {
	return &AuthService{Repo: repo, JWTSecret: secret, TokenTTL: tokenTTL}
}

func (s *AuthService) Register(ctx context.Context, login, password string) (string, string, error) {
	if len(login) < 3 || len(password) < 6 {
		return "", "", errors.New("login or password too short")
	}

	exists, err := s.Repo.ExistsByLogin(ctx, login)
	if err != nil {
		return "", "", err
	}
	if exists {
		return "", "", errors.New("user already exists")
	}

	hash, err := auth.HashPassword(password)
	if err != nil {
		return "", "", err
	}

	userID := uuid.New().String()
	if err := s.Repo.CreateUser(ctx, userID, login, hash); err != nil {
		return "", "", err
	}

	return userID, login, nil
}

func (s *AuthService) Login(ctx context.Context, login, password string) (string, error) {
	user, err := s.Repo.GetByLogin(ctx, login)
	if err != nil {
		return "", err
	}

	if err := auth.ComparePassword(user.PasswordHash, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := auth.GenerateToken(user.ID, user.Login, s.JWTSecret, s.TokenTTL)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, token string) (*auth.Claims, error) {
	return auth.ValidateToken(token, s.JWTSecret)
}
