package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value;
	fmt.Println(<-ch);
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return t1.Left == t2.Left;
}

func main() {
	ch := make(chan int, 1);

	go Walk(tree.New(1), ch);

	// for index := 0; index < 10; index++ {
	// 	select {
	// 	case <-ch: 
	// 		fmt.Println(<-ch)
	// 	}
	// }
}
