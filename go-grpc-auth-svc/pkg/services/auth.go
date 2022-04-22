package services

import (
	"context"
	"fmt"
	"github.com/karakaya/go-grpc-project/go-grpc-auth-svc/pkg/db"
	"github.com/karakaya/go-grpc-project/go-grpc-auth-svc/pkg/models"
	"github.com/karakaya/go-grpc-project/go-grpc-auth-svc/pkg/pb"
	"github.com/karakaya/go-grpc-project/go-grpc-auth-svc/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	result := s.H.DB.Database("usersdb").Collection("users").FindOne(ctx, bson.M{"email": req.GetEmail()})
	fmt.Println(result.Err())
	if result.Err() != mongo.ErrNoDocuments {
		return &pb.RegisterResponse{Status: http.StatusConflict, Error: "This email already exists"}, nil
	}

	user.Email = req.GetEmail()
	user.Password = utils.HashPassword(req.GetPassword())

	s.H.DB.Database("usersdb").Collection("users	").InsertOne(ctx, bson.M{"email": user.Email, "password": user.Password})
	return &pb.RegisterResponse{Status: http.StatusCreated}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	if result := s.H.DB.Database("usersdb").Collection("users").FindOne(ctx, bson.M{"email": req.Email}); result.Err() != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var user models.User
	if result := s.H.DB.Database("usersdb").Collection("users").FindOne(ctx, bson.M{"email": claims.Email}); result.Err() != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "user not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil

}
