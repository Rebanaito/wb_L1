package main

import (
	"fmt"
	"sync"
)

// Function that calculates the squared value of a number.
// The result is sent to the summing function through a channel.
func calcSquare(n int, ch chan<- int, wg *sync.WaitGroup) {
	n *= n
	ch <- n
}

// Function that listens on the channel for squared values.
// Each value is incremented to the total sum.
// Mutex is used to avoid problems with concurrent routines
// writing to a variable at the same time.
func changeSum(ch <-chan int, sum *int, wg *sync.WaitGroup) {
	mu := sync.Mutex{}
	for square := range ch {
		mu.Lock()
		*sum += square
		mu.Unlock()
		wg.Done()
	}
}

func main() {
	array := []int{2, 4, 6, 8, 10}

	// Channel for communication between routines
	ch := make(chan int)

	// Wait group to make sure all values are added first
	// before printing the sum
	sum := 0
	wg := &sync.WaitGroup{}

	// Launching the sum function as a routine
	go changeSum(ch, &sum, wg)

	// Each number of the array is calculated concurrently
	for _, num := range array {
		wg.Add(1)
		go calcSquare(num, ch, wg)
	}

	// Waiting for all routines to finish before printing
	// the result
	wg.Wait()
	fmt.Print(sum)
}
