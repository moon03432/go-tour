package main

import (
	"fmt"
	"math"
)

// const
const Pi = 3.14

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// uint(x): type conversion
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
