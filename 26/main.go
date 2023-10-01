package main

import "fmt"

// This function uses an array of bool values big enough
// to accomodate all ASCII values. All uppercase letters
// are marked as lowercase to make the function case-insensitive.
func allSymbolsAreUnique(str string) bool {
	chars := make([]bool, 128)
	for _, letter := range str {
		index := int(letter)
		if index >= 65 && index <= 90 {
			index += 32
		}
		if chars[int(letter)] {
			return false
		} else {
			chars[int(letter)] = true
		}
	}
	return true
}

func main() {
	fmt.Println("abcd - ", allSymbolsAreUnique("abcd"))
	fmt.Println("abCdefAaf - ", allSymbolsAreUnique("abCdefAaf"))
	fmt.Println("aabcd - ", allSymbolsAreUnique("aabcd"))
}
