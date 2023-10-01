package main

import "fmt"

// This function finds numbers that are present in both arrays.
// It achieves this by mapping the values of the shorter array
// and then iterating over the longer one to find any matches
func findIntersecting(shorter, longer []int) []int {
	var intersect []int = nil
	nums := make(map[int]struct{})
	for _, num := range shorter {
		nums[num] = struct{}{}
	}
	for _, num := range longer {
		_, ok := nums[num]
		if ok {
			if intersect == nil {
				intersect = make([]int, 1)
				intersect[0] = num
			} else {
				intersect = append(intersect, num)
			}
		}
	}
	return intersect
}

// This function simply checks which array is shorter
func intersection(one, two []int) []int {
	if len(one) < len(two) {
		return findIntersecting(one, two)
	}
	return findIntersecting(two, one)
}

func main() {
	one := []int{24, 5, -7, 2, 10, 26, 33, -42}
	two := []int{-7, 25, 57, 26, 11, 4}
	intersect := intersection(one, two)
	if intersect == nil {
		fmt.Println("No intersection between arrays")
	} else {
		fmt.Println(intersect)
	}
}
