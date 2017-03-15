package main

import (
	"fmt"
	"math/cmplx"
)

// var: declaration
var c, python, java bool

// outside function, every statement begins with a keyword (var, func, etc). so := doesn't work
//k := 3

// good way for declaring variables
var (
	ToBe  bool = false
	MaxInt uint64 = 1<<64-1
	z complex128 = cmplx.Sqrt(-5+12i)
)

func main() {
	var i,j int = 1,2
	//c, python, java = true, false, "no!"

	// type inference (推断)
	k := 3
	fmt.Println(i, j, k, c, python, java)

	fmt.Println(ToBe, MaxInt, z)
}