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

// The worker function that will call the increment method on the counter.
// We have to use mutex to make sure concurrent writes to a variable don't
// break anything, while wait group makes sure we don't miss an increment
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
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go routine(&counter, i%5+1, &mu, wg)
	}
	wg.Wait()
	fmt.Println(counter)
}
