package db

import (
	"context"
	"github.com/karakaya/go-grpc-project/go-grpc-auth-svc/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Handler struct {
	DB *mongo.Client
}

func ConnectDB(c *config.Config) Handler {
	dbCredentials := options.Credential{
		Username: "root",
		Password: "example",
	}
	clientOptions := options.Client().ApplyURI(c.HOST).SetAuth(dbCredentials)

	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return Handler{DB: client}
}
