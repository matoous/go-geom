package transform

import (
	"math"

	"github.com/twpayne/go-geom"
)

const (
	x_PI = 3.14159265358979324 * 3000.0 / 180.0
)

// UniqueCoords creates a new coordinate array (with the same layout as the inputs) that
// contains each unique coordinate in the coordData.  The ordering of the coords are the
// same as the input.
func UniqueCoords(layout geom.Layout, compare Compare, coordData []float64) []float64 {
	set := NewTreeSet(layout, compare)
	stride := layout.Stride()
	uniqueCoords := make([]float64, 0, len(coordData))
	numCoordsAdded := 0
	for i := 0; i < len(coordData); i += stride {
		coord := coordData[i : i+stride]
		added := set.Insert(geom.Coord(coord))

		if added {
			uniqueCoords = append(uniqueCoords, coord...)
			numCoordsAdded++
		}
	}
	return uniqueCoords[:numCoordsAdded*stride]
}

func BD09ToGCJ02(x, y float64) (float64, float64) {
	tmpX := x - 0.0065
	tmpY := y - 0.006
	z := math.Sqrt(tmpX*tmpX+tmpY*tmpY) - 0.00002*math.Sin(tmpY*x_PI)
	theta := math.Atan2(tmpY, tmpX) - 0.000003*math.Cos(tmpX*x_PI)
	retX := z * math.Cos(theta)
	retY := z * math.Sin(theta)
	return retX, retY
}

func GCJ02ToBD09(x, y float64) (float64, float64) {
	z := math.Sqrt(x*x+y*y) + 0.00002*math.Sin(y*x_PI)
	theta := math.Atan2(y, x) + 0.000003*math.Cos(x*x_PI)
	retX := z*math.Cos(theta) + 0.0065
	retY := z*math.Sin(theta) + 0.006
	return retX, retY
}

func WGS84ToBD09(x, y float64) (float64, float64) {
	return GCJ02ToBD09(WGS84ToGCJ02(x, y))
}

func BD09ToWGS84(x, y float64) (float64, float64) {
	return GCJ02ToWGS84(BD09ToGCJ02(x, y))
}
