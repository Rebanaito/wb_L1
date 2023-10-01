package main

import (
	"fmt"
	"sync"
)

// Function that calculates the squared value of a number.
// The result is sent to the output function through a channel.
func calcSquare(n int, ch chan int, wg *sync.WaitGroup) {
	n *= n
	ch <- n
}

// Function that handles the output to stdout
func output(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Little flag for formatting reasons
	first := true

	// Listening on the channel and printing each received value
	for square := range ch {
		if first {
			first = false
			fmt.Print(square)
		} else {
			fmt.Print(" ", square)
		}
	}
}

func main() {
	array := []int{2, 4, 6, 8, 10}

	// Channel for the calculation routine to communicate with output.
	// Waitgroup is passed to the output function to make sure all values
	// are printed before exiting the program.
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	// Launching the output function as a routine
	go output(ch, wg)

	// Launching concurrent goroutines to calculate the squared values
	for _, num := range array {
		wg.Add(1)
		go calcSquare(num, ch, wg)
	}

	// Waiting for the final result to be printed to the screen
	wg.Wait()
}
