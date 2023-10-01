package main

import "fmt"

// Function that sets the 'i'th bit to 'value'
// 'Value' is 1 or 0, 'i' is a number between 0 and 64
// The target bit is accessed through a mask variable and
// bitwise operators
func changebit(num *int64, i uint, value bool) {
	// Create a mask - a variable where all bits are 0
	// except for bit 'i' (chosen by the user).
	mask := int64(1) << (i - 1)

	if value {
		// Bitwise OR to set the 'i' bit to 1
		*num |= mask
	} else {
		// Bitwise AND NOT (bitclear) to set the bit 'i' to 0
		*num &^= mask
	}
}

func main() {
	var num int64
	var i uint
	var value bool
	fmt.Print("What number do you want to work with?: ")
	fmt.Scan(&num)
	fmt.Print("Which bit do you want to change?: ")
	fmt.Scan(&i)
	if i > 64 {
		fmt.Print("i has to be >= 1 && <= 64")
		return
	}
	fmt.Print("What value should the 'i' bit be set to?: ")
	fmt.Scan(&value)
	fmt.Printf("Old value: %64b\n", num)
	changebit(&num, i, value)
	fmt.Printf("New value: %64b", num)
}
