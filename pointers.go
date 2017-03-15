package main

import "fmt"

/*
 * Go has no pointer arithmetic, unlike C
 */
func main() {
	i, j := 42, 2701

	// pointer zero value is nil
	var p *int = nil

	p = &i

	// "*": dereference / indirection operator
	fmt.Println(*p)  // 42
	*p = 21
	fmt.Println(i)   // 21

	p = &j           // point to j
	*p = *p / 37
	fmt.Println(j)   // 73
}