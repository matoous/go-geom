package geom

import (
	"math"
)

// Distance calculates the distance between two [Coord] in a desired unit.
// This uses the [Haversine formula](http://en.wikipedia.org/wiki/Haversine_formula) to account for global curvature.
func Distance(from Coord, to Coord) Meters {
	var dLat = DegreesToRadians(Degrees(to[1] - from[1]))
	var dLon = DegreesToRadians(Degrees(to[0] - from[0]))
	var lat1 = DegreesToRadians(Degrees(from[1]))
	var lat2 = DegreesToRadians(Degrees(to[1]))

	var a = math.Pow(math.Sin(float64(dLat)/2), 2) +
		math.Pow(math.Sin(float64(dLon)/2), 2)*math.Cos(float64(lat1))*math.Cos(float64(lat2))

	return RadiansToLength(Radians(2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))))
}
