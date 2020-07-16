package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"go-grpc/location"
)

func main() {
	listen, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	location.RegisterLocationServer(grpcServer, &location.Server{})

	log.Println("Starting gRPC server...")

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
