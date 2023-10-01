package main

import "fmt"

// From the little context we have in our example, there is
// no apparent reason to use a global variable. Global variables
// tend to lead to unexpected behavior and bugs, and their use
// should be avoided unless absolutely necessary

// Here we assume that the number passed to this function
// specifies the lenght of the created string...
func createHugeString(len int) string {
	str := make([]rune, len)
	return string(str)
}

// ...however, it's probably not the best practice to perform
// operations on memory with hardcoded values like this, at
// least not without checking if it would cause any problems,
// which is what we do here
func someFunc() string {
	v := createHugeString(1 << 10)
	if len(v) > 100 {
		return v[:100]
	}
	return v
}

func main() {
	justString := someFunc()
	fmt.Println(len(justString))
}
