package main

import "fmt"

// A dog 'class' that has its own method
type dog struct {
}

func (d *dog) bark() {
	fmt.Println("WOOF")
}

// Same for cat
type cat struct {
}

func (c *cat) makeSound() {
	fmt.Println("MEOW")
}

// We are a dog person and don't
// know how to interact with cats.
// However, we get a small book that
// teaches us how invoke similar
// responses from a cat using our knowledge
// about dogs
type catAdapter struct {
	cat *cat
}

// A cat has his own weird way of barking, but to
// us it's a bark nonetheless
func (a *catAdapter) bark() {
	a.cat.makeSound()
}

// This struct will now allow us to interact with both
// pets the same way. Since we are a dog person,
// we will generally approach our cat the same way as a dog,
// just with some adaptions
type animal struct {
}

type pet interface {
	bark()
}

func (pet *animal) makeSound(p pet) {
	p.bark()
}

func main() {
	// We decide to get a pet
	pet := &animal{}

	// It's a dog!
	dog := &dog{}

	// We learned how to make our pet make sounds
	// It barks!
	pet.makeSound(dog)

	// Now there's also a cat
	cat := &cat{}

	// We don't know what a cat is, but apparently
	// it's not too different from a dog. We simply
	// need to adapt our approach a bit
	adapter := &catAdapter{cat}

	// Nice! Our new pet can make sounds too
	// It has a strange way of barking...
	pet.makeSound(adapter)
}
