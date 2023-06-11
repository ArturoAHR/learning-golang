package main

// https://go.dev/tour/moretypes/23

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	counter := make(map[string]int)
	for _, word := range strings.Fields(s) {
		counter[word]++
	}
	return counter
}

func main() {
	wc.Test(WordCount)
}
