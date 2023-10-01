package main

import (
	"fmt"
	"time"
)

// Our sleep function marks the time when it is called
// and sits in an infinite loop until the time passed since
// the start is equal to the duration provided
func sleep(duration time.Duration) {
	start := time.Now()
	for time.Since(start) < duration {
	}
}

func main() {
	fmt.Print("Somebody ")
	sleep(3 * time.Second)
	fmt.Print("once told me...")
}
