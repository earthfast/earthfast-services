package geo

import "math"

const (
	// Radius of the Earth in meters
	earthRadius = 6378100.0
)

// Distance calculates the distance (in meters) between two Coordinates, using the
// Haversine formula: http://en.wikipedia.org/wiki/Haversine_formula.
//
// Adapted from https://gist.github.com/cdipaolo/d3f8db3848278b49db68
func Distance(loc1, loc2 Coordinate) float64 {
	// Convert to radians
	var la1, lo1, la2, lo2 float64
	la1 = loc1.Latitude * math.Pi / 180
	lo1 = loc1.Longitude * math.Pi / 180
	la2 = loc2.Latitude * math.Pi / 180
	lo2 = loc2.Longitude * math.Pi / 180

	// Calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * earthRadius * math.Asin(math.Sqrt(h))
}

// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
