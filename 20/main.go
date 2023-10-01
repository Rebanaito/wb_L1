package main

import (
	"fmt"
	"strings"
)

// We turn the string into a string array using the Split function.
// The rest is the same as mirroring a word in the last task
func reverseStrings(str string) string {
	words := strings.Split(str, " ")
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	newString := strings.Join(words, " ")
	return newString
}

func main() {
	str := reverseStrings("snow dog sun")
	fmt.Println(str)
}
