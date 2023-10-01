package main

import (
	"fmt"
	"reflect"
)

// We can infer the type of a variable through an interface
// by using the TypeOf() function from the standard Go library
// package 'reflect'
func main() {
	var a, b, c, d, e, f interface{}
	var ch chan interface{}
	a = 10
	b = 10.4
	c = "string"
	d = true
	e = ch
	fmt.Print(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c), reflect.TypeOf(d), reflect.TypeOf(e), reflect.TypeOf(f))
}
