package main

import "fmt"

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
