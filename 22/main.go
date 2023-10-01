package main

import (
	"fmt"
	"math/big"
)

func main() {
	var A, B string
	var a, b big.Int
	fmt.Print("Please provide two numbers: ")
	fmt.Scan(&A, &B)
	a.SetString(A, 10)
	b.SetString(B, 10)
	var sum, sub, div, mul big.Int
	sum.Add(&a, &b)
	sub.Sub(&a, &b)
	div.Div(&a, &b)
	mul.Mul(&a, &b)
	fmt.Println("a = ", a, ", b = ", b)
	fmt.Println("a + b = ", sum)
	fmt.Println("a - b = ", sub)
	fmt.Println("a / b = ", div)
	fmt.Println("a * b = ", mul)
}
