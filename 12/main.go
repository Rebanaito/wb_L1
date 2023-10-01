package main

import "fmt"

// Since Go doesn't have sets, we can improvise one by using
// a map where the value is a an empty struct interface
func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})
	for _, word := range strings {
		set[word] = struct{}{}
	}
	fmt.Println(set)
}
