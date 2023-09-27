package main

import (
	"distance/point"
	"fmt"
	"math"
)

func distance(A, B *point.Point) float64 {
	return math.Sqrt(math.Pow(float64(B.GetX()-A.GetX()), 2) + math.Pow(float64(B.GetY()-A.GetY()), 2))
}

func main() {
	A := point.NewPoint(-2, 5)
	B := point.NewPoint(3, -10)
	fmt.Println(distance(A, B))
}
