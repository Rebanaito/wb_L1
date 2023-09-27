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

func timer(N int, wg *sync.WaitGroup) {

}

func worker(ctx context.Context, id, N int, data chan int, wg *sync.WaitGroup) {
	for {
		select {
		case num := <-data:
			fmt.Println("Worker ", id, " took job ", num)
			//time.Sleep(time.Duration(N) * time.Second)
		case <-ctx.Done():
			fmt.Println("Worker ", id, " is done for today")
			wg.Done()
			return
		}
	}
}

func main() {
	var N int
	fmt.Printf("Provide N - number of workers: ")
	fmt.Scan(&N)
	if N < 1 {
		fmt.Printf("N must be > 0")
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	data := make(chan int)
	wg.Add(N)
	for i := 0; i < N; i++ {
		go worker(ctx, i+1, N, data, wg)
	}
	go func() {
		for i := 1; ; i++ {
			data <- i
			time.Sleep(time.Second)
		}
	}()
	kill := make(chan os.Signal)
	signal.Notify(kill, syscall.SIGINT, syscall.SIGTERM)
	<-kill
	cancel()
	wg.Wait()
	fmt.Println("All workers done, closing")
}
