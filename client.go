package main

import (
	"google.golang.org/grpc"
	"log"
	"go-grpc/location"
	"context"
	"fmt"
)

func main() {
	// create grpc connection
	connection, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer connection.Close()

	// create grpc client
	client := location.NewLocationClient(connection)

	// send request
	request := location.DistanceRequest{
		Point1: &location.Point{Latitude: 50.450001, Longitude: 30.523333}, // Kiev
		Point2: &location.Point{Latitude: 49.842957, Longitude: 24.031111}, // Lviv
	}

	response, err := client.Distance(context.Background(), &request)
	if err != nil {
		log.Fatalf("GRPC request error: %s", err)
	}

	fmt.Println("GRPC response:", response.Distance)
}
