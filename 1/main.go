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

func (h *Human) SayName() {
	fmt.Printf("My name is %s\n", h.Name)
}

func (h *Human) SayAge() {
	fmt.Printf("I am %d years old\n", h.Age)
}

type Person struct {
	Human
	Gender string
}

func (p *Person) SayGender() {
	fmt.Printf("I am a %s\n", p.Gender)
}

func main() {
	person := Person{Human{}, "male"}
	person.Name = "Bilbo Baggins"
	person.Age = 19
	person.SayName()
	person.SayAge()
	person.SayGender()
}
