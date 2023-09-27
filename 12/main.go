package main

import "fmt"

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})
	for _, word := range strings {
		set[word] = struct{}{}
	}
	fmt.Println(set)
}
