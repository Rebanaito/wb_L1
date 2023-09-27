package main

import (
	"fmt"
	"sync"
)

func writeToMap(dict map[string]int, word string, index int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	dict[word] = index
	mu.Unlock()
}

func main() {
	words := []string{"Somebody", "Once", "Told", "Me", "The", "World", "Is", "Gonna", "Roll"}
	dict := make(map[string]int)
	wg := &sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(len(words))
	for i, word := range words {
		go writeToMap(dict, word, i, wg, &mu)
	}
	wg.Wait()
	fmt.Println(dict)
}
