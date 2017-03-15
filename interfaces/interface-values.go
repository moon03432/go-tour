package main

import (
	"fmt"
	"math"
)

/*
 * interface contains (value, type)
 */
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {

	// nil interface
	// calling a method to nil interface is a run-time error
	var i I
	describe(i)

	// interface with nil value
	// interface (nil, type) is non-nil
	var t *T
	i = t
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
