package client

import (
	"github.com/Babushkin05/simple-marketplace/api-gateway/internal/config"
	"google.golang.org/grpc"
)

type Clients struct {
	Auth  AuthClient
	Goods GoodsClient
}

func InitClients(cfg config.Config) (*Clients, error) {
	authConn, err := grpc.Dial(cfg.AuthService.Addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	goodsConn, err := grpc.Dial(cfg.GoodsService.Addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Clients{
		Auth:  NewAuthClient(authConn),
		Goods: NewGoodsClient(goodsConn),
	}, nil
}
