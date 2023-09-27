package point

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func (p *Point) GetX() int {
	return p.x
}

func (p *Point) GetY() int {
	return p.y
}
