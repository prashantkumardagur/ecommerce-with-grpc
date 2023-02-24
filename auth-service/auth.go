package main

import (
	"context"

	pb "github.com/prashantkumardagur/ecommerce-with-grpc/proto"
)

//==================================================================================================

var User = pb.User{
	Id:       "1",
	Name:     "Prashant Kumar",
	Email:    "prashantkumardagur@gmail.com",
	Password: "123456",
	Phone:    "1234567890",
}

//==================================================================================================

func (s *server) Ping(ctx context.Context, req *pb.Empty) (*pb.Response, error) {
	return &pb.Response{Success: true, Message: "Pong"}, nil
}

//==================================================================================================

func (s *server) GetProfile(ctx context.Context, req *pb.Id) (*pb.User, error) {
	return &User, nil
}

//==================================================================================================

func (s *server) Register(ctx context.Context, req *pb.User) (*pb.Response, error) {
	User = *req
	return &pb.Response{Success: true, Message: "User registered successfully"}, nil
}
