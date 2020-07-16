package location

import (
	"log"
	"golang.org/x/net/context"
	"math"
)

type Server struct {

}

func (server *Server) Distance(context context.Context, request *DistanceRequest) (*DistanceResponse, error) {
	log.Println("Received distance GRPC request:")
	log.Printf("-- Point 1: %s", request.Point1)
	log.Printf("-- Point 2: %s", request.Point2)

	distance := Distance(request.Point1.Latitude, request.Point1.Longitude, request.Point2.Latitude, request.Point2.Longitude)

	return &DistanceResponse{Distance: distance}, nil
}

// Distance function returns the distance (in meters) between two points of
// a given longitude and latitude relatively accurately (using a spherical
// approximation of the Earth) through the Haversin Distance Formula for
// great arc distance on a sphere with accuracy for small distances
//
// http://en.wikipedia.org/wiki/Haversine_formula
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	const R float64 = 6371000 // Earth radius

	// convert to radians
	var f1 = lat1 * math.Pi / 180
	var f2 = lat2 * math.Pi / 180
	var l1 = lon1 * math.Pi / 180
	var l2 = lon2 * math.Pi / 180

	// calculate
	h := math.Pow(math.Sin((f2 - f1) / 2), 2) + math.Cos(f1) * math.Cos(f2) * math.Pow(math.Sin((l2 - l1) / 2), 2)
	d := 2 * R * math.Asin(math.Sqrt(h))

	return d
}
