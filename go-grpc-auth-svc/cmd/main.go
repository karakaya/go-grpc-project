package main

import (
	"github.com/karakaya/go-grpc-project/go-grpc-auth-svc/pkg/config"
	"github.com/karakaya/go-grpc-project/go-grpc-auth-svc/pkg/db"
	"github.com/karakaya/go-grpc-project/go-grpc-auth-svc/pkg/pb"
	"github.com/karakaya/go-grpc-project/go-grpc-auth-svc/pkg/services"
	"github.com/karakaya/go-grpc-project/go-grpc-auth-svc/pkg/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	//c, err := config.LoadConfig()
	//if err != nil {
	//	log.Fatalln("failed at config", err)
	//}

	c := config.Config{
		PORT:         ":27017",
		AppPort:      ":50051",
		HOST:         "mongodb://localhost:27017/",
		PASSWORD:     "86w3792tHela0iR4",
		USERNAME:     "doadmin",
		DATABASE:     "usersdb",
		PROTOCOL:     "mongodb+srv",
		JWTSecretKey: "r43t18sc",
	}

	h := db.ConnectDB(&c)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.AppPort)
	if err != nil {
		log.Fatalln("failed to listen", err)
	}
	log.Println("listening at: ", c.AppPort)

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("failed to server: ", err)
	}
}
