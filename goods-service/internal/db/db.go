package db

import "github.com/Babushkin05/simple-marketplace/goods-service/internal/models"

type GoodsRepository interface {
	CreateAd(ad *models.Ad) (*models.Ad, error)
	ListAds(filter models.AdsFilter) ([]*models.AdWithAuthor, error)
}
