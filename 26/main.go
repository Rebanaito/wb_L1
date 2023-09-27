package main

import "fmt"

func allLettersAreUnique(str string) bool {
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
	fmt.Println("abcd - ", allLettersAreUnique("abcd"))
	fmt.Println("abCdefAaf - ", allLettersAreUnique("abCdefAaf"))
	fmt.Println("aabcd - ", allLettersAreUnique("aabcd"))
}
