package main

import "fmt"

// We can avoid creating a temporary variable to swap
// two variables by using this feature of the Go language.
// (even though is probably still has to create one under the hood...)
func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 21
	b := 42
	fmt.Printf("a: %d, b: %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("a: %d, b: %d", a, b)
}
