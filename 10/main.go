package main

import (
	"fmt"
)

// Function that adds a number from the array to the map.
// If the group for the respective temperature group already
// exists, it appends to the group, otherwise it creates a new group
func insertNum(num float32, ranges map[int][]float32) {
	group := (int(num) / 10) * 10
	_, ok := ranges[group]
	if !ok {
		ranges[group] = *new([]float32)
	}
	ranges[group] = append(ranges[group], num)
}

func main() {
	// Map that stores all temperature values grouped by ranges
	ranges := make(map[int][]float32)

	temps := []float32{-25.4, -27.0, 13.0, -7.3, 5.6, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Adding numbers to the map iteratively, using routines doesn't seem
	// to speed anything up since all routines get stuck in a mutex
	// queue
	for _, num := range temps {
		insertNum(num, ranges)
	}
	fmt.Println(ranges)
}
