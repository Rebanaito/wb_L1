package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// We turn the string into a string array using the Split function.
// The rest is the same as mirroring a word in the last task
func reverseSentence(str string) string {
	words := strings.Split(str, " ")
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	newString := strings.Join(words, " ")
	return newString
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(`What sentence would you like to mirror?: `)
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]
	str := reverseSentence(input)
	fmt.Println(str)
}
