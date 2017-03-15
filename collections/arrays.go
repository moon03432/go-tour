package main

import (
	"fmt"
	"strings"
)

/*
 * arrays & slices
 * slice is reference to array, which doesn't store data, but is more commonly used
 */
func main() {
	var a1 = [2]string {
		"hello",
		"world",
	}
	fmt.Println(a1)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// []int (slice) is like reference to array, doesn't store data
	var s []int = primes[1:4]	// [3 5 7]
	fmt.Println(s)

	primes[2] = 6
	fmt.Println(s)

	// slice to array of customized struct
	s2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s2)

	// defaults values
	s3 := []int{2, 3, 5, 7, 11, 13}
	s3 = s3[1:4]		// [3 5 7]
	printSlice(s3)		// len=3 cap=5
	s3 = s3[:2]		// [3 5]
	printSlice(s3)		// len=2 cap=5
	s3 = s3[1:]		// [5]
	printSlice(s3)		// len=1 cap=4, as slice is a reference, thus moving forward will reduce capacity

	// nil is the default value of a slice
	var s4 []int
	printSlice(s4)		// nil

	// make(type, length, capacity) -> slice
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)

	// Slice of slices
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func printSlice(s []int) {
	if s == nil {
		fmt.Println("nil")
	} else {
		// len(slice) & cap(slice): capacity
		fmt.Printf("%v len=%d cap=%d\n", s, len(s), cap(s))
	}
}