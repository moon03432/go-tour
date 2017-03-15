package main

import "fmt"

/*
 * for-range: iterates over a slice / map
 */
var pow = []int{1,2,4,8,16,32,64,128}

func main(){

	for i,v := range pow  {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// use "_" to omit index
	for _, value := range pow {
		fmt.Printf("%d ", value)
	}
}
