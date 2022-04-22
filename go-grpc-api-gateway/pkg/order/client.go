package order

import (
	"fmt"

	"github.com/karakaya/go-grpc-project/go-grpc-api-gateway/pkg/config"
	"github.com/karakaya/go-grpc-project/go-grpc-api-gateway/pkg/order/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect:", err)
	}
	return pb.NewOrderServiceClient(cc)
}
