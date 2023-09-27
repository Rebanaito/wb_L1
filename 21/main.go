package main

import "fmt"

type pet interface {
	bark()
}

type dog struct {
}

func (d *dog) bark() {
	fmt.Println("WOOF")
}

type cat struct {
}

func (c *cat) meow() {
	fmt.Println("MEOW")
}

type adapter struct {
	cat *cat
}

func (a *adapter) bark() {
	a.cat.meow()
}

type animal struct {
}

func (pet *animal) makeSound(p pet) {
	p.bark()
}

func main() {
	pet := &animal{}
	dog := &dog{}
	pet.makeSound(dog)
	cat := &cat{}
	adapter := &adapter{cat}
	pet.makeSound(adapter)
}
