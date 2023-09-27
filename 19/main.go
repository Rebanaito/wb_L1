package main

import "fmt"

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
