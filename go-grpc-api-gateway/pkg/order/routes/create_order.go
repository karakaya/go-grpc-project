package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karakaya/go-grpc-project/go-grpc-api-gateway/pkg/order/pb"
)

type CreateOrderRequestBody struct {
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	body := CreateOrderRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	userId, _ := ctx.Get("userId")
	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		UserId:    userId.(int64),
		Quantity:  body.Quantity,
		ProductId: body.ProductId,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusCreated, &res)
}
