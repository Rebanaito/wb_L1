package main

import "fmt"

// We turn the string into a rune array and iterate until the middle
// of the array, swapping mirroring elements - first with last, second
// with second to last and so on...
func reverseString(str string) string {
	chars := []rune(str)
	for i := 0; i < len(chars)/2; i++ {
		chars[i], chars[len(chars)-1-i] = chars[len(chars)-1-i], chars[i]
	}
	return string(chars)
}

func main() {
	str := reverseString("reverseString")
	fmt.Println(str)
}
