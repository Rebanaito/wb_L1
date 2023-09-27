package main

import "fmt"

func changebit(num *int64, i int, value bool) {
	mask := int64(1) << (i - 1)
	if value {
		*num |= mask
	} else {
		*num &^= mask
	}
}

func main() {
	var num int64
	var i int
	var value bool
	fmt.Print("What number do you want to work with?: ")
	fmt.Scan(&num)
	fmt.Print("Which bit do you want to change?: ")
	fmt.Scan(&i)
	if i < 1 || i > 64 {
		fmt.Print("i has to be >= 1 && <= 64")
		return
	}
	fmt.Print("What value should the 'i' bit be set to?: ")
	fmt.Scan(&value)
	fmt.Printf("Old value: %64b\n", num)
	changebit(&num, i, value)
	fmt.Printf("New value: %64b", num)
}
