package main

import (
	"fmt"
	"strings"
	"time"
)

func join(str []string) {
	start := time.Now()
	defer fmt.Println("Join took ", time.Since(start))
	new := strings.Join(str, "")
	fmt.Println(new)
}

func plus(str []string) {
	start := time.Now()
	defer fmt.Println("Plus took ", time.Since(start))
	var new string
	for _, word := range str {
		new += word
	}
	fmt.Println(new)
}

func builder(str []string) {
	start := time.Now()
	defer fmt.Println("Builder took ", time.Since(start))
	builder := strings.Builder{}
	for _, word := range str {
		builder.WriteString(word)
	}
	fmt.Println(builder.String())
}

func sprintf(str []string) {
	start := time.Now()
	defer fmt.Println("Sprintf took ", time.Since(start))
	new := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", str[0], str[1], str[2], str[3], str[4], str[5], str[6], str[7], str[8], str[9])
	fmt.Println(new)
}

func main() {
	strings := []string{"Somebody", "once", "told", "me", "the", "world", "is", "gonna", "roll", "me"}
	join(strings)
	plus(strings)
	builder(strings)
	sprintf(strings)
}
