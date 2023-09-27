package main

import (
	"fmt"
	"sync"
)

func calcSquare(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	n *= n
	ch <- n
}

func main() {
	array := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	first := true
	go func() {
		for square := range ch {
			if first {
				first = false
				fmt.Print(square)
			} else {
				fmt.Print(" ", square)
			}
		}
	}()
	wg := &sync.WaitGroup{}
	for _, num := range array {
		wg.Add(1)
		go calcSquare(num, ch, wg)
	}
	wg.Wait()
}
