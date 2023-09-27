package main

import "fmt"

func intersection(one, two []int) []int {
	var intersect []int = nil
	nums := make(map[int]struct{})
	if len(one) < len(two) {
		for _, num := range one {
			nums[num] = struct{}{}
		}
		for _, num := range two {
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
	} else {
		for _, num := range two {
			nums[num] = struct{}{}
		}
		for _, num := range one {
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
	}
	return intersect
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
