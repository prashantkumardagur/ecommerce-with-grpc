package main

import (
	"context"

	pb "github.com/prashantkumardagur/ecommerce-with-grpc/proto"
)

//==================================================================================================

var Catalog = pb.Catalog{
	Products: []*pb.Product{},
}

//==================================================================================================

func (s *server) Ping(ctx context.Context, req *pb.Empty) (*pb.Response, error) {
	return &pb.Response{Success: true, Message: "Pong"}, nil
}

//==================================================================================================

func (s *server) GetCatalog(ctx context.Context, req *pb.Empty) (*pb.Catalog, error) {
	return &Catalog, nil
}

//==================================================================================================

func (s *server) AddProduct(ctx context.Context, req *pb.Product) (*pb.Response, error) {
	Catalog.Products = append(Catalog.Products, req)
	return &pb.Response{Success: true, Message: "Product added to catalog successfully"}, nil
}

//==================================================================================================

func (s *server) DeleteProduct(ctx context.Context, req *pb.Id) (*pb.Response, error) {
	for i, product := range Catalog.Products {
		if product.Id == req.Id {
			Catalog.Products = append(Catalog.Products[:i], Catalog.Products[i+1:]...)
			return &pb.Response{Success: true, Message: "Product deleted from catalog successfully"}, nil
		}
	}
	return &pb.Response{Success: false, Message: "Product not found in catalog"}, nil
}

//==================================================================================================

func (s *server) UpdateProduct(ctx context.Context, req *pb.Product) (*pb.Response, error) {
	for i, product := range Catalog.Products {
		if product.Id == req.Id {
			Catalog.Products[i] = req
			return &pb.Response{Success: true, Message: "Product updated in catalog successfully"}, nil
		}
	}
	return &pb.Response{Success: false, Message: "Product not found in catalog"}, nil
}
