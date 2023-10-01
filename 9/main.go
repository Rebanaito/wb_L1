package main

import (
	"fmt"
	"sync"
)

// Output function reads from the second channel and
// writes to stdout
func output(outChan chan int, wg *sync.WaitGroup) {
	for num := range outChan {
		fmt.Printf("%d ", num)
		wg.Done()
	}
}

// Input function reads from the first channel and
// writes the doubled value to the second channel
func input(inChan, outChan chan int) {
	for num := range inChan {
		outChan <- num * 2
	}
}

func main() {
	// Two channels - first takes values from the array,
	// second one passes modified values to the output function
	inChan := make(chan int)
	outChan := make(chan int)
	array := []int{10, 7, 245, 63, 2, -19, 25, 44}

	// Wait group to make sure all values are printed first
	// before exit
	wg := &sync.WaitGroup{}

	// Starting input and output functions as routines
	go output(outChan, wg)
	go input(inChan, outChan)

	// Passing all numbers from the array to the first channel
	for _, num := range array {
		wg.Add(1)
		inChan <- num
	}

	// Waiting for all modified values to be printed
	wg.Wait()
}
