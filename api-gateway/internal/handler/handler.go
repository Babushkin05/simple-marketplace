package handler

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	authpb "github.com/Babushkin05/simple-marketplace/api-gateway/api/gen/auth"
	goodspb "github.com/Babushkin05/simple-marketplace/api-gateway/api/gen/goods"
)

type Handler struct {
	authClient  authpb.AuthServiceClient
	goodsClient goodspb.GoodsServiceClient
}

func NewHandler(authConn, goodsConn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()

	h := &Handler{
		authClient:  authpb.NewAuthServiceClient(authConn),
		goodsClient: goodspb.NewGoodsServiceClient(goodsConn),
	}

	router.POST("/register", h.Register)
	router.POST("/login", h.Login)

	router.POST("/ads", h.CreateAd)
	router.GET("/ads", h.ListAds)

	return router
}
