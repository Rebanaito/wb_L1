package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var N int
	fmt.Println("How long should the program run?:")
	fmt.Scan(&N)
	if N < 1 {
		fmt.Println("N should be > 0")
		return
	}
	ch := make(chan string)
	wg := &sync.WaitGroup{}
	wg.Add(N)
	go func() {
		for message := range ch {
			fmt.Println(message)
			wg.Done()
		}
	}()
	for i := 1; i <= N; i++ {
		time.Sleep(time.Second)
		ch <- fmt.Sprintf("Message #%d", i)
	}
	wg.Wait()
}
