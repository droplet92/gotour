package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		if t.Left != nil {
			Walk(t.Left, ch)
		}
		ch <- t.Value
		if t.Right != nil {
			Walk(t.Right, ch)
		}
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		a, b := <-ch1, <-ch2
		fmt.Printf("compares: %d vs %d\n", a, b)
		if a != b {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(
		Same(tree.New(1), tree.New(1)),
		Same(tree.New(1), tree.New(2)),
	)
}
