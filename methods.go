package main

import (
	"fmt"
	"math"
)

// A method on type is a function with "receiver"
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// This is a regular function
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// Receiver is a pointer: avoid copy the value on each method call
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

// method definition with a receiver whose type should be defined in the same package as the method
// so methods on float64 should be defined in the built-in package
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3,4}
	v.Scale(10)
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}


