package point

// A point struct with unexported fields x and y
type Point struct {
	x int
	y int
}

// Constructor function that lets us create a new point with set values
func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

// Two exported methods that allow us to get values of the point
func (p *Point) GetX() int {
	return p.x
}

func (p *Point) GetY() int {
	return p.y
}
