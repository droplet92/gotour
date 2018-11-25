package main

import "fmt"

func fibonacci() func() int {
	pprev, prev := 0, 1
	tmp := 0
	return func() int {
		tmp = pprev
		pprev, prev = prev, pprev+prev
		return tmp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
