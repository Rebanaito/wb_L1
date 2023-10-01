package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// The worker function.
// Select allows us to listen for the channel AND the canceling of the context
// at the same time, handling whichever comes first.
// When a number is received on the main channel, worker prints it to stdout.
// When the context is canceled, the worker notifies about
// it and finishes gracefully
func worker(ctx context.Context, id, N int, data chan int, wg *sync.WaitGroup) {
	for {
		select {
		case num := <-data:
			fmt.Println("Worker ", id, " took job ", num)
		case <-ctx.Done():
			fmt.Println("Worker ", id, " is done for today")
			wg.Done()
			return
		}
	}
}

// Function that keeps sending data to the main channel indefinitely
func mainChannel(data chan int) {
	for i := 1; ; i++ {
		data <- i
		time.Sleep(time.Second)
	}
}

// Function for the graceful handling of the interrupt signal
func interrupt() {
	kill := make(chan os.Signal)
	signal.Notify(kill, syscall.SIGINT, syscall.SIGTERM)
	<-kill
}

func main() {
	var N int
	fmt.Printf("Provide N - number of workers: ")
	fmt.Scan(&N)
	if N < 1 {
		fmt.Fprintf(os.Stderr, "N must be > 0")
		return
	}

	// The context variable that lets us handle the interrupt signal gracefully.
	// Cancel function terminates the context, forcing workers to stop.
	ctx, cancel := context.WithCancel(context.Background())

	// Wait group helps us make sure all workers have stopped before finishing.
	wg := &sync.WaitGroup{}

	// The main data channel
	data := make(chan int)

	// Launch the number of workers specified by the user
	for i := 0; i < N; i++ {
		wg.Add(1)
		go worker(ctx, i+1, N, data, wg)
	}

	// Start sending data to the channel
	go mainChannel(data)

	// Wait for the interrupt signal
	interrupt()

	// Cancel the context and wait for all workers to finish
	cancel()
	wg.Wait()
	fmt.Println("All workers done, closing")
}
