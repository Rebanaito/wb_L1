package main

import "fmt"

func quicksort(array []int, low, high int) {
	if low >= high {
		return
	}
	pivot := quicksortPartition(array, low, high)
	quicksort(array, low, pivot-1)
	quicksort(array, pivot+1, high)
}

func quicksortPartition(array []int, low, high int) int {
	pivot := array[high]
	i := low - 1
	for j := low; j < high; j++ {
		if array[j] < pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[high] = array[high], array[i+1]
	return i + 1
}

func main() {
	array := []int{10, -2, 4, 1, 26, -20, 17, -18}
	fmt.Println("Before sort: ", array)
	quicksort(array, 0, len(array)-1)
	fmt.Println("After sort: ", array)
}
