package square

import "math"

// Define custom int type to hold sides number.
type SidesFigure int

// Define constants to represent 0, 3 and 4 sides. Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0).
const (
	SidesCircle   SidesFigure = 0
	SidesTriangle SidesFigure = 3
	SidesSquare   SidesFigure = 4
)

// Define triangle factor as internal (constant) variable.
var triangleFactor float64 = math.Sqrt(3) / 4

func CalcSquare(sideLen float64, sidesNum SidesFigure) float64 {
	switch sidesNum {
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)
	case SidesTriangle:
		return triangleFactor * math.Pow(sideLen, 2)
	case SidesSquare:
		return math.Pow(sideLen, 2)
	}
	return 0
}
