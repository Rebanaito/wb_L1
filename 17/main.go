package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		middle := left + (right-left)/2
		if arr[middle] == target {
			return middle
		}
		if arr[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 6, 7, 9, 14, 25, 26, 45, 46, 51, 63, 79}
	fmt.Println(binarySearch(arr, 25))
}
