package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// This routine waits for a message on the channel signaling it to stop.
// Select is used to have a default action that is performed until the
// signal arrives
func routineOne(ch <-chan int, trackChan chan<- int, wg *sync.WaitGroup) {
	for {
		select {
		case <-ch:
			trackChan <- 0
			fmt.Println("Routine one finished")
			wg.Done()
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

// This routine keeps reading messages from the channel and
// stops as soon as the channel is closed
func routineTwo(ch <-chan string, trackChan chan<- int, wg *sync.WaitGroup) {
	for {
		_, ok := <-ch
		if !ok {
			trackChan <- 1
			fmt.Println("Routine two done")
			wg.Done()
			return
		}
	}
}

// This routine keeps running until the context it listens to is canceled
func routineThree(ctx context.Context, trackChan chan<- int, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			trackChan <- 2
			fmt.Println("Routine three done")
			wg.Done()
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func keepTrack(trackChan <-chan int) {
	routines := []bool{true, true, true}
	for {
		select {
		case value := <-trackChan:
			routines[value] = false
		default:
			fmt.Print("Routines running:")
			for i := range routines {
				if routines[i] {
					fmt.Print(" ", i+1)
				}
			}
			fmt.Println()
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// Context for routine #3
	ctx, cancel := context.WithCancel(context.Background())

	// Wait group to make sure all routines are done before exit
	wg := &sync.WaitGroup{}

	// Channels for routines #1 and #2 and the channel to track all channels
	channelOne := make(chan int)
	channelTwo := make(chan string)
	trackChan := make(chan int)

	// Starting all three routines
	wg.Add(3)
	go routineOne(channelOne, trackChan, wg)
	go routineTwo(channelTwo, trackChan, wg)
	go routineThree(ctx, trackChan, wg)
	go keepTrack(trackChan)

	// Sending data to channel #2
	go func() {
		for i := 0; i < 6; i++ {
			channelTwo <- "blank"
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
