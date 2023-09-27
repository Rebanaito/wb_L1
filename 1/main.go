package main

import "fmt"

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
