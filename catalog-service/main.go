package main

import (
	"context"
	"log"
	"net"

	pb "github.com/prashantkumardagur/ecommerce-with-grpc/proto"

	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//==================================================================================================

type server struct {
	pb.CatalogServiceServer
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

var Collection *mongo.Collection

//==================================================================================================

func main() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://prashantkumar:Password024680@testcluster.8xzqf.mongodb.net/?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, e := mongo.Connect(context.TODO(), clientOptions)
	if e != nil {
		return
	}

	// Check the connection
	e = client.Ping(context.TODO(), nil)
	if e != nil {
		return
	}

	// get collection as ref
	Collection = client.Database("ecommerce").Collection("catalog")

	//==================================================================================================

	// Create a tcp listener
	lis, tcperr := net.Listen("tcp", ":50053")
	HandleError(tcperr)

	// Create a gRPC server
	grpcServer := grpc.NewServer()

	// Register catalog service on the gRPC server
	pb.RegisterCatalogServiceServer(grpcServer, &server{})

	// Serve the gRPC server on the tcp listener
	log.Println("Server is running on port 50053")
	grpcErr := grpcServer.Serve(lis)
	HandleError(grpcErr)
}
