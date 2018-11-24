package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	words := strings.Fields(s)

	for _, value := range(words) {
		_, ok := res[value]
		if ok == true {
			res[value]++
		} else {
			res[value] = 1
		}
	}
	return res
}

func main() {
	wc.Test(WordCount)
}