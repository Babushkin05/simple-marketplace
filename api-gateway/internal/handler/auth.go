package handler

import (
	"context"
	"net/http"

	authpb "github.com/Babushkin05/simple-marketplace/api-gateway/api/gen/auth"
	"github.com/gin-gonic/gin"
)

type registerRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register godoc
// @Summary      Зарегистрировать нового пользователя
// @Description  Создаёт нового пользователя по логину и паролю
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body handler.registerRequest true "Данные регистрации"
// @Success      200 {object} map[string]string "Успешный ответ с user_id и login"
// @Failure      400 {object} map[string]string "Некорректный ввод"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /register [post]
func (h *Handler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.authClient.Register(context.Background(), &authpb.RegisterRequest{
		Login:    req.Login,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": resp.UserId, "login": resp.Login})
}

type loginRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login godoc
// @Summary      Авторизация пользователя
// @Description  Авторизует пользователя и возвращает JWT токен
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body handler.loginRequest true "Данные авторизации"
// @Success      200 {object} map[string]string "Успешный ответ с токеном"
// @Failure      400 {object} map[string]string "Некорректный ввод"
// @Failure      401 {object} map[string]string "Неверные логин или пароль"
// @Router       /login [post]
func (h *Handler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.authClient.Login(context.Background(), &authpb.LoginRequest{
		Login:    req.Login,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": resp.Token})
}
