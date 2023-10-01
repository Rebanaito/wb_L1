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
	if b.String() != "0" {
		div.Div(&a, &b)
	}
	mul.Mul(&a, &b)
	division := div.String()
	if division == "0" {
		division = "Cannot divide by zero"
	}
	fmt.Println("a = ", a, ", b = ", b)
	fmt.Println("a + b = ", sum.String())
	fmt.Println("a - b = ", sub.String())
	fmt.Println("a / b = ", division)
	fmt.Println("a * b = ", mul.String())
}
