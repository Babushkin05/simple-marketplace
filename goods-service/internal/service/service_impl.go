package service

import (
	"context"
	"errors"
	"time"

	authpb "github.com/Babushkin05/simple-marketplace/goods-service/api/gen/auth"
	"github.com/Babushkin05/simple-marketplace/goods-service/internal/db"
	"github.com/Babushkin05/simple-marketplace/goods-service/internal/models"
)

var (
	ErrInvalidToken    = errors.New("invalid token")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrValidationError = errors.New("validation error")
)

type service struct {
	repo db.GoodsRepository
	auth authpb.AuthServiceClient
}

func NewService(repo db.GoodsRepository, auth authpb.AuthServiceClient) GoodsService {
	return &service{
		repo: repo,
		auth: auth,
	}
}

func (s *service) CreateAd(ctx context.Context, token string, ad models.Ad) (models.Ad, error) {
	resp, err := s.auth.ValidateToken(ctx, &authpb.ValidateTokenRequest{Token: token})

	if err != nil || resp == nil {
		return models.Ad{}, ErrUnauthorized
	}

	if len(ad.Title) == 0 || len(ad.Title) > 100 {
		return models.Ad{}, ErrValidationError
	}
	if len(ad.Description) == 0 || len(ad.Description) > 1000 {
		return models.Ad{}, ErrValidationError
	}
	if len(ad.ImageURL) == 0 {
		return models.Ad{}, ErrValidationError
	}
	if ad.Price <= 0 {
		return models.Ad{}, ErrValidationError
	}

	ad.AuthorID = resp.Login
	ad.CreatedAt = time.Now()

	res, err := s.repo.CreateAd(&ad)
	if err != nil {
		return models.Ad{}, err
	}

	return *res, nil
}

func (s *service) GetAds(ctx context.Context, token string, filter models.AdsFilter) ([]models.AdWithAuthor, error) {
	var userID string

	if token != "" {
		resp, err := s.auth.ValidateToken(ctx, &authpb.ValidateTokenRequest{Token: token})
		if err != nil || resp == nil {
			userID = ""
		}
	}

	ads, err := s.repo.ListAds(filter)
	if err != nil {
		return nil, err
	}

	result := make([]models.AdWithAuthor, 0, len(ads))
	for _, ad := range ads {
		isMine := userID != "" && ad.AuthorID == userID
		result = append(result, models.AdWithAuthor{
			Ad:     *ad,
			IsMine: isMine,
		})
	}

	return result, nil
}
