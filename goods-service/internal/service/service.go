package service

import (
	"context"

	"github.com/Babushkin05/simple-marketplace/goods-service/internal/models"
)

type AuthClient interface {
	ValidateToken(ctx context.Context, token string) (userID string, login string, valid bool, err error)
}

type GoodsService interface {
	CreateAd(ctx context.Context, token string, ad models.Ad) (models.Ad, error)
	GetAds(ctx context.Context, token string, filter models.AdsFilter) ([]models.AdWithAuthor, error)
}
