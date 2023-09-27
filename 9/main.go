package main

import (
	"fmt"
	"sync"
)

func main() {
	inChan := make(chan int)
	outChan := make(chan int)
	array := []int{10, 7, 245, 63, 2, -19, 25, 44}
	wg := &sync.WaitGroup{}
	go func() {
		for num := range outChan {
			fmt.Printf("%d ", num)
			wg.Done()
		}
	}()
	go func() {
		for num := range inChan {
			outChan <- num * 2
		}
	}()
	wg.Add(len(array))
	for _, num := range array {
		inChan <- num
	}
	wg.Wait()
}
