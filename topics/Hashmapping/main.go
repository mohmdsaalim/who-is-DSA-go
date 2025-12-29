package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go is fast go is simple"
    fmt.Println(WordFrequency(text))
}

func WordFrequency(s string) map[string]int {
	freq := make(map[string]int)
	words := strings.Fields(s)
	fmt.Println(words)
	for _,v := range words{
		freq[v]++
	}
	return freq
}