package main

import (
	"fmt"
	"math"
)

/*
https://tour.golang.org/methods/20
 */

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	var a float64 = 10

	for i:=0;i<10 ;i++  {
		a -= (a*a-x)/a/2
	}

	return a, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(-2))
}