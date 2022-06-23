package main

import (
	"google.golang.org/grpc"
	"log"
	"go-grpc/location"
	"context"
	"fmt"
)

type City struct {
	name string
	latitude float64
	longitude float64
}

func main() {
	// create grpc connection
	connection, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer connection.Close()

	// create grpc client
	client := location.NewLocationClient(connection)

	cities := []City{
		{"Lviv", 49.842957, 24.031111},
		{"Kiev", 50.450001, 30.523333},
	}

	position := &location.Point{Latitude: 48.5522, Longitude: 24.4238} // IF

	// send requests
	for _, city := range cities {
		request := location.DistanceRequest{
			Point1: position,
			Point2: &location.Point{Latitude: city.latitude, Longitude: city.longitude},
		}

		response, err := client.Distance(context.Background(), &request)
		if err != nil {
			log.Fatalf("GRPC request error: %s", err)
		}

		fmt.Printf("Distance to %s: %f\n", city.name, response.Distance)
	}
}
