package main

import "fmt"

/* func append(slice, elements...) -> new slice
 */

func main() {
	var s []int
	PrintSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	PrintSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	PrintSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	PrintSlice(s)
}

func PrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
