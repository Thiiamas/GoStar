package utils

//Point ...
type (
	Point struct {
		X, Y int
	}
)

func newPoint(x,y int) Point {
	return &Point{
		X: x
		Y: y
	}

	

	
}
