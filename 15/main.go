package main

import "fmt"

var justString string

func createHugeString(len int) string {
	fmt.Println("Len is ", len)
	str := make([]rune, len)
	return string(str)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
	fmt.Println(len(justString))
}
