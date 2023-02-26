package main

import (
	"context"

	pb "github.com/prashantkumardagur/ecommerce-with-grpc/proto"
	"go.mongodb.org/mongo-driver/bson"
)

//==================================================================================================

func (s *server) Ping(ctx context.Context, req *pb.Empty) (*pb.Response, error) {
	return &pb.Response{Success: true, Message: "Pong"}, nil
}

//==================================================================================================

func (s *server) GetCatalog(ctx context.Context, req *pb.Empty) (*pb.Catalog, error) {
	var Catalog pb.Catalog

	// Get all products from database
	dbCatalog, err := Collection.Find(context.Background(), bson.M{})
	HandleError(err)

	// Decode all products and append to Catalog
	for dbCatalog.Next(context.Background()) {
		var product pb.Product
		err := dbCatalog.Decode(&product)
		HandleError(err)
		Catalog.Products = append(Catalog.Products, &product)
	}

	return &Catalog, nil
}

//==================================================================================================

func (s *server) GetProduct(ctx context.Context, req *pb.Id) (*pb.Product, error) {
	// Get product from database
	dbProduct := Collection.FindOne(context.Background(), bson.M{"id": req.Id})
	if dbProduct.Err() != nil {
		return nil, dbProduct.Err()
	}

	// Decode product
	var product pb.Product
	err := dbProduct.Decode(&product)
	HandleError(err)

	return &product, nil
}

//==================================================================================================

func (s *server) AddProduct(ctx context.Context, req *pb.Product) (*pb.Response, error) {
	// Check if product already exists
	dbProduct := Collection.FindOne(context.Background(), bson.M{"id": req.Id})
	if dbProduct.Err() == nil {
		return &pb.Response{Success: false, Message: "Product already exists"}, nil
	}

	// Add product to database
	_, err := Collection.InsertOne(context.Background(), req)
	HandleError(err)

	return &pb.Response{Success: true, Message: "Product added to catalog successfully"}, nil
}

//==================================================================================================

func (s *server) DeleteProduct(ctx context.Context, req *pb.Id) (*pb.Response, error) {
	// Delete product from database
	_, err := Collection.DeleteMany(context.Background(), bson.M{"id": req.Id})
	if err != nil {
		return &pb.Response{Success: false, Message: "Product not found"}, nil
	}

	return &pb.Response{Success: true, Message: "Product deleted from catalog successfully"}, nil
}

//==================================================================================================

func (s *server) UpdateProduct(ctx context.Context, req *pb.Product) (*pb.Response, error) {
	// Update product in database
	_, err := Collection.UpdateOne(context.Background(), bson.M{"id": req.Id}, bson.M{"$set": req})
	if err != nil {
		return &pb.Response{Success: false, Message: "Product not found"}, nil
	}

	return &pb.Response{Success: true, Message: "Product updated in catalog successfully"}, nil
}
