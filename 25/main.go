package main

import (
	"fmt"
	"time"
)

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
