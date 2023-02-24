package main

import (
	"net"

	pb "github.com/prashantkumardagur/ecommerce-with-grpc/proto"

	"google.golang.org/grpc"
)

//==================================================================================================

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

type server struct {
	pb.AuthServiceServer
}

//==================================================================================================

func main() {

	// Create a TCP listener
	lis, tcpErr := net.Listen("tcp", ":50051")
	HandleErr(tcpErr)

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach auth service to the server
	pb.RegisterAuthServiceServer(grpcServer, &server{})

	// start the server
	err := grpcServer.Serve(lis)
	HandleErr(err)
}
