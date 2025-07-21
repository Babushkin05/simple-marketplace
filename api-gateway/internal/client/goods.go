package client

import (
	goodspb "github.com/Babushkin05/simple-marketplace/api-gateway/api/gen/goods"
	"google.golang.org/grpc"
)

type GoodsClient interface {
	goodspb.GoodsServiceClient
}

func NewGoodsClient(conn *grpc.ClientConn) GoodsClient {
	return goodspb.NewGoodsServiceClient(conn)
}
