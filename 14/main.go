package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a, b, c, d, e interface{}
	var ch chan interface{}
	a = 10
	b = 10.4
	c = "string"
	d = true
	e = ch
	fmt.Print(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c), reflect.TypeOf(d), reflect.TypeOf(e))
}
