package main

import "fmt"

// Since Go is not an OOP language, it does not have proper inheritance.
// However, we can simulate inheritance by embedding structs within structs.
// If we pretend that struct is a class, a child 'class' (struct) simply has
// to have its parent 'classes' (structs) embedded in order to use their
// methods.

type Human struct {
	Name string
	Age  int
}

func (h Human) SayName() {
	fmt.Printf("My name is %s\n", h.Name)
}

func (h Human) SayAge() {
	fmt.Printf("I am %d years old\n", h.Age)
}

type Action struct {
	Human
}

func main() {
	act := Action{}
	act.Name = "Bilbo Baggins"
	act.Age = 19
	act.SayName()
	act.SayAge()
}
