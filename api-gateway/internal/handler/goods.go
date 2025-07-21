package handler

import (
	"context"
	"net/http"
	"strconv"

	goodspb "github.com/Babushkin05/simple-marketplace/api-gateway/api/gen/goods"
	"github.com/gin-gonic/gin"
)

type createAdRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	ImageURL    string  `json:"image_url" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

// CreateAd godoc
// @Summary      Создать объявление
// @Description  Создаёт новое объявление. Требуется JWT токен в заголовке.
// @Tags         ads
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "JWT токен"
// @Param        input body createAdRequest true "Данные объявления"
// @Success      200 {object} goodspb.CreateAdResponse "Успешно созданное объявление"
// @Failure      400 {object} map[string]string "Некорректный запрос"
// @Failure      401 {object} map[string]string "Нет токена авторизации"
// @Failure      500 {object} map[string]string "Ошибка сервера"
// @Router       /ads [post]
func (h *Handler) CreateAd(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
		return
	}

	var req createAdRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	adResp, err := h.goodsClient.CreateAd(context.Background(), &goodspb.CreateAdRequest{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageURL,
		Price:       req.Price,
		Token:       token,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, adResp)
}

// ListAds godoc
// @Summary      Получить список объявлений
// @Description  Получает список объявлений с возможностью сортировки и фильтрации. JWT токен (опционально) для персонализации.
// @Tags         ads
// @Accept       json
// @Produce      json
// @Param        Authorization header string false "JWT токен"
// @Param        page query int false "Номер страницы" default(1)
// @Param        page_size query int false "Размер страницы" default(10)
// @Param        min_price query number false "Минимальная цена"
// @Param        max_price query number false "Максимальная цена"
// @Param        sort_by query string false "Поле сортировки (created_at, price)" default(created_at)
// @Param        sort_order query string false "Порядок сортировки (asc, desc)" default(asc)
// @Success      200 {array} goodspb.CreateAdResponse "Список объявлений"
// @Failure      500 {object} map[string]string "Ошибка сервера"
// @Security     BearerAuth
// @Router       /ads [get]
func (h *Handler) ListAds(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	sortByStr := c.DefaultQuery("sort_by", "created_at")
	sortOrderStr := c.DefaultQuery("sort_order", "asc")

	minPriceStr := c.Query("min_price")
	maxPriceStr := c.Query("max_price")

	var sortBy goodspb.ListAdsRequest_SortField
	switch sortByStr {
	case "price":
		sortBy = goodspb.ListAdsRequest_PRICE
	case "created_at":
		sortBy = goodspb.ListAdsRequest_CREATED_AT
	default:
		sortBy = goodspb.ListAdsRequest_CREATED_AT
	}

	var sortOrder goodspb.ListAdsRequest_SortOrder
	switch sortOrderStr {
	case "desc":
		sortOrder = goodspb.ListAdsRequest_DESC
	case "asc":
		sortOrder = goodspb.ListAdsRequest_ASC
	default:
		sortOrder = goodspb.ListAdsRequest_ASC
	}

	var priceMin, priceMax float64
	if v, err := strconv.ParseFloat(minPriceStr, 64); err == nil {
		priceMin = v
	}
	if v, err := strconv.ParseFloat(maxPriceStr, 64); err == nil {
		priceMax = v
	}

	req := &goodspb.ListAdsRequest{
		Page:      int32(page),
		PageSize:  int32(pageSize),
		SortBy:    sortBy,
		SortOrder: sortOrder,
		PriceMin:  priceMin,
		PriceMax:  priceMax,
	}

	resp, err := h.goodsClient.ListAds(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.Ads)
}
