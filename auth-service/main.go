package main

import (
	"context"
	"log"
	"net"

	pb "github.com/prashantkumardagur/ecommerce-with-grpc/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc"
)

//==================================================================================================

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

type server struct {
	pb.AuthServiceServer
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
	Collection = client.Database("ecommerce").Collection("users")

	//==================================================================================================

	// Create a TCP listener
	lis, tcpErr := net.Listen("tcp", ":50051")
	HandleError(tcpErr)

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach auth service to the server
	pb.RegisterAuthServiceServer(grpcServer, &server{})

	// start the server
	log.Println("Server is running on port 50051")
	err := grpcServer.Serve(lis)
	HandleError(err)
}
