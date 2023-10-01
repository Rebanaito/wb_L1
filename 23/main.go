package main

import "fmt"

func removeElement(slice []int, index int) []int {
	index -= 1
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)
	slice = removeElement(slice, 5)
	fmt.Println(slice)
}
