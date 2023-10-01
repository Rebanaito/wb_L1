package main

import (
	"fmt"
	"sync"
	"time"
)

// Listener function that reads from the channel
func listener(ch chan string, wg *sync.WaitGroup) {
	for message := range ch {
		fmt.Println(message)
		wg.Done()
	}
}

func main() {
	var N int
	fmt.Println("How long should the program run?:")
	fmt.Scan(&N)
	if N < 1 {
		fmt.Println("N should be > 0")
		return
	}

	// The main data channel
	ch := make(chan string)

	// Wait group to make sure all messages are received before exit
	wg := &sync.WaitGroup{}

	// Starting listener
	go listener(ch, wg)

	// Sending data to the channel for the duration specified by user
	for i := 1; i <= N; i++ {
		wg.Add(1)
		time.Sleep(time.Second)
		ch <- fmt.Sprintf("Message #%d", i)
	}

	// Waiting for the listener to print the last message
	wg.Wait()
}
