package main

import (
	"fmt"
	"sync"
)

func calcSquare(n int, ch chan int, wg *sync.WaitGroup) {
	n *= n
	ch <- n
}

func main() {
	array := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	mu := sync.Mutex{}
	sum := 0
	wg := &sync.WaitGroup{}
	wg.Add(len(array))
	go func() {
		for square := range ch {
			mu.Lock()
			sum += square
			mu.Unlock()
			wg.Done()
		}
	}()
	for _, num := range array {
		go calcSquare(num, ch, wg)
	}
	wg.Wait()
	fmt.Print(sum)
}
