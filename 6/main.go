package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func routineOne(ch chan int, wg *sync.WaitGroup) {
	for {
		select {
		case <-ch:
			fmt.Println("Routine one finished")
			wg.Done()
			return
		default:
			fmt.Println("Routine one still running")
			time.Sleep(1 * time.Second)
		}
	}
}

func routineTwo(ch chan string, wg *sync.WaitGroup) {
	for {
		message, ok := <-ch
		if !ok {
			fmt.Println("Routine two done")
			wg.Done()
			return
		}
		fmt.Println(message)
	}
}

func routineThree(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Routine three done")
			wg.Done()
			return
		default:
			fmt.Println("Routine three still running")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	channelOne := make(chan int)
	channelTwo := make(chan string)
	wg.Add(3)
	go routineOne(channelOne, wg)
	go routineTwo(channelTwo, wg)
	go routineThree(ctx, wg)
	go func() {
		for i := 0; i < 6; i++ {
			channelTwo <- "Routine two still running"
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("Sending stop signal to routine one")
	channelOne <- 1
	time.Sleep(3 * time.Second)
	fmt.Println("Closing channel two to shut down routine two")
	close(channelTwo)
	time.Sleep(3 * time.Second)
	fmt.Println("Cancelling the context to shut down routine three")
	cancel()
	wg.Wait()
	fmt.Println("All routines have finished")
}
