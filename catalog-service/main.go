package main

import (
	"net"

	pb "github.com/prashantkumardagur/ecommerce-with-grpc/proto"

	"google.golang.org/grpc"
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

//==================================================================================================

func main() {
	// Create a tcp listener
	lis, tcperr := net.Listen("tcp", ":50053")
	HandleError(tcperr)

	// Create a gRPC server
	grpcServer := grpc.NewServer()

	// Register catalog service on the gRPC server
	pb.RegisterCatalogServiceServer(grpcServer, &server{})

	// Serve the gRPC server on the tcp listener
	grpcErr := grpcServer.Serve(lis)
	HandleError(grpcErr)
}