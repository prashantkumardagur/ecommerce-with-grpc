package main

import (
	"context"

	pb "github.com/prashantkumardagur/ecommerce-with-grpc/proto"
)

//==================================================================================================

var Cart = pb.Cart{
	UserId:   "1",
	Products: []*pb.Product{},
}

//==================================================================================================

func (s *server) Ping(ctx context.Context, req *pb.Empty) (*pb.Response, error) {
	return &pb.Response{Success: true, Message: "Pong"}, nil
}

//==================================================================================================

func (s *server) GetCart(ctx context.Context, req *pb.Id) (*pb.Cart, error) {
	return &Cart, nil
}

//==================================================================================================

func (s *server) AddToCart(ctx context.Context, req *pb.Id) (*pb.Response, error) {
	Cart.Products = append(Cart.Products, &pb.Product{
		Id:          req.Id,
		Name:        "Product " + req.Id,
		Price:       100,
		Description: "This is a product with id " + req.Id,
		Image:       "/some/image/path",
	})
	return &pb.Response{Success: true, Message: "Product added to cart successfully"}, nil
}
