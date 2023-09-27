package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	Count uint64
}

func (c *Counter) Increment() {
	c.Count++
}

func routine(counter *Counter, delay int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(delay))
	mu.Lock()
	counter.Increment()
	mu.Unlock()
}

func main() {
	var counter Counter
	mu := sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go routine(&counter, i%5+1, &mu, wg)
	}
	wg.Wait()
	fmt.Println(counter)
}
