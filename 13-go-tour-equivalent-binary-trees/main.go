package main

// https://go.dev/tour/concurrency/8

import (
	"fmt"

	"golang.org/x/tour/tree"
)

/*
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
*/

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		value1, value2 := <-ch1, <-ch2
		if value1 != value2 {
			return false
		}
	}
	return true
}

func main() {
	channel := make(chan int)
	go Walk(tree.New(1), channel)

	for i := 0; i < 10; i++ {
		value := <-channel
		fmt.Println(value)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
}
