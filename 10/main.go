package main

import (
	"fmt"
)

func insertNum(num float32, ranges map[int][]float32) {
	group := int(num) / 10
	group *= 10
	_, ok := ranges[group]
	if ok {
		ranges[group] = append(ranges[group], num)
	} else {
		ranges[group] = make([]float32, 1)
		ranges[group][0] = num
	}
}

func main() {
	ranges := make(map[int][]float32)
	temps := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for _, num := range temps {
		insertNum(num, ranges)
	}
	fmt.Println(ranges)
}
