package main

/*

 There can be many different binary trees with the same sequence of values stored at the leaves.
 For example, here are two binary trees storing the sequence 1, 1, 2, 3, 5, 8, 13.

                 3                    8
               /  \                  / \
              1   8                 3  13
             / \ / \               / \
            1  2 5 15             1  5
                                 / \
                                1  2

 (In this example, right child's value > a node's value > left child's value)

 A function to check whether two binary trees store the same sequence is quite complex in most languages.
 We'll use Go's concurrency and channels to write a simple solution.

 This example uses the tree package, which defines the type:

 type Tree struct {
     Left  *Tree
     Value int
     Right *Tree
 }

 1. Implement the Walk function.

 2. Test the Walk function.

    The function tree.New(k) constructs a randomly-structured binary tree holding the values k, 2k, 3k, ..., 10k.

    Create a new channel ch and kick off the walker:

        go Walk(tree.New(1), ch)

    Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

 3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.

 4. Test the Same function.

    Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

 */

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkTree(t, ch)
	close(ch)
}

func WalkTree(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	WalkTree(t.Left, ch)
	ch <- t.Value
	WalkTree(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1,ch2 := make(chan int), make(chan int)

	go Walk(t1,ch1)
	go Walk(t2,ch2)

	for {
		v1, ok1 := <- ch1
		v2, ok2 := <- ch2

		if ok1 != ok2 || v1 != v2 {
			return false
		}

		if !ok1 {
			break
		}
	}

	return true
}

func main() {
	ch := make(chan int)

	go Walk(tree.New(1), ch)

	for v := range ch {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}