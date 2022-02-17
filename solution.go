package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type ShapeSides int

const (
	SidesCircle ShapeSides = 0
	SidesTriangle = 3
	SidesSquare = 4
)

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum ShapeSides) float64 {
	switch sidesNum {
	case SidesCircle:
		// should be (math.Pi * sideLen * sideLen)
		return math.Pi * sideLen * sideLen / 2.0
	case SidesTriangle:
		// should be (sideLen * sidelen * math.Sqrt(3.0) / 4.0)
		return sideLen * math.Sqrt(3.0) / 4.0
	case SidesSquare:
		return sideLen * sideLen
	default:
		return 0.0
	}
}
