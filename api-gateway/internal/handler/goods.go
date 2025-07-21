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
// @Description  Получает список объявлений с пагинацией. Требуется JWT токен в заголовке.
// @Tags         ads
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "JWT токен"
// @Param        page query int false "Номер страницы" default(1)
// @Param        page_size query int false "Размер страницы" default(10)
// @Success      200 {array} goodspb.CreateAdResponse "Список объявлений"
// @Failure      401 {object} map[string]string "Нет токена авторизации"
// @Failure      500 {object} map[string]string "Ошибка сервера"
// @Router       /ads [get]
func (h *Handler) ListAds(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
		return
	}

	// Параметры пагинации
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	resp, err := h.goodsClient.ListAds(context.Background(), &goodspb.ListAdsRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Token:    token,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.Ads)
}
