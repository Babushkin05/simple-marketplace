package grpc

import (
	"context"
	"errors"
	"strings"
	"time"

	authpb "github.com/Babushkin05/simple-marketplace/goods-service/api/gen/auth"
	goodspb "github.com/Babushkin05/simple-marketplace/goods-service/api/gen/goods"
	"github.com/Babushkin05/simple-marketplace/goods-service/internal/models"
	"github.com/Babushkin05/simple-marketplace/goods-service/internal/service"
)

type GoodsHandler struct {
	goodspb.UnimplementedGoodsServiceServer
	svc         service.GoodsService
	authService authpb.AuthServiceClient
}

func NewGoodsHandler(svc service.GoodsService, authClient authpb.AuthServiceClient) *GoodsHandler {
	return &GoodsHandler{
		svc:         svc,
		authService: authClient,
	}
}

func (h *GoodsHandler) CreateAd(ctx context.Context, req *goodspb.CreateAdRequest) (*goodspb.CreateAdResponse, error) {
	// Валидация токена
	authResp, err := h.authService.ValidateToken(ctx, &authpb.ValidateTokenRequest{
		Token: req.Token,
	})
	if err != nil || authResp.Login == "" {
		return nil, errors.New("unauthorized")
	}

	ad := models.Ad{
		Title:       req.Title,
		Description: req.Description,
		ImageURL:    req.ImageUrl,
		Price:       req.Price,
		AuthorID:    authResp.UserId,
		CreatedAt:   time.Now(),
	}

	created, err := h.svc.CreateAd(ctx, req.Token, ad)
	if err != nil {
		return nil, err
	}

	return &goodspb.CreateAdResponse{
		Id:          created.ID,
		Title:       created.Title,
		Description: created.Description,
		ImageUrl:    created.ImageURL,
		Price:       created.Price,
	}, nil
}

func (h *GoodsHandler) ListAds(ctx context.Context, req *goodspb.ListAdsRequest) (*goodspb.ListAdsResponse, error) {
	filter := models.AdsFilter{
		MinPrice: &req.PriceMin,
		MaxPrice: &req.PriceMax,
		SortBy:   strings.ToLower(req.SortBy.String()),
		SortDesc: req.SortOrder == goodspb.ListAdsRequest_DESC,
		Limit:    int(req.PageSize),
		Offset:   (int(req.Page) - 1) * int(req.PageSize),
	}

	ads, err := h.svc.GetAds(ctx, req.Token, filter)
	if err != nil {
		return nil, err
	}
	var result []*goodspb.CreateAdResponse
	for _, ad := range ads {
		result = append(result, &goodspb.CreateAdResponse{
			Id:          ad.ID,
			Title:       ad.Title,
			Description: ad.Description,
			ImageUrl:    ad.ImageURL,
			Price:       ad.Price,
			AuthorLogin: ad.AuthorID,
			IsOwner:     ad.IsMine,
		})
	}

	return &goodspb.ListAdsResponse{Ads: result}, nil
}
