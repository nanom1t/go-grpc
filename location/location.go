package location

import (
	"log"
	"golang.org/x/net/context"
)

type Server struct {

}

func (server *Server) Distance(context context.Context, request *DistanceRequest) (*DistanceResponse, error) {
	log.Println("Received distance GRPC request:")
	log.Printf("-- Point 1: %s", request.Point1)
	log.Printf("-- Point 2: %s", request.Point2)

	var distance float64 = 25

	return &DistanceResponse{Distance: distance}, nil
}
