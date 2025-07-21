package service

import (
	"context"

	"github.com/Babushkin05/simple-marketplace/goods-service/internal/models"
)

type GoodsService interface {
	CreateAd(ctx context.Context, token string, ad models.Ad) (models.Ad, error)
	GetAds(ctx context.Context, token string, filter models.AdsFilter) ([]models.AdWithAuthor, error)
}
