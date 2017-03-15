package main

import (
	"fmt"
)

/*
 * switch v := i.(type) {
 * case T:
 *    // here v has type T
 * case S:
 *    // here v has type S
 * default:
 *    // no match; here v has the same type as i
 * }
 */

func do(i interface{})  {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v * 2 = %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)

	}
}

func main()  {
	do(21)
	do("hello")
	do(true)
}
