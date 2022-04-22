package main

import (
	"github.com/gin-gonic/gin"
	"github.com/karakaya/go-grpc-project/go-grpc-api-gateway/pkg/auth"
	"github.com/karakaya/go-grpc-project/go-grpc-api-gateway/pkg/config"
	"github.com/karakaya/go-grpc-project/go-grpc-api-gateway/pkg/order"
	"github.com/karakaya/go-grpc-project/go-grpc-api-gateway/pkg/product"
)

func main() {
	//c, err := config.LoadConfig()
	//if err != nil {
	//	log.Fatalln("failed to load config", err)
	//}

	c := config.Config{
		Port:          ":3000",
		AuthSvcUrl:    "localhost:50051",
		ProductSvcUrl: "localhost:50052",
		OrderSvcUrl:   "localhost:50053",
	}
	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)
	r.Run(c.Port)
}
