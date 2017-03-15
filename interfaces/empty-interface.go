package main

import (
	"fmt"
)

func main()  {

	// Empty interface may hold values of any type
	// useful to handle values of unknown type
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}