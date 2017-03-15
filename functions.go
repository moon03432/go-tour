package main

import "fmt"

func add(x int, y int) int {
	return x+y
}

func substract(x,y int) int {
	return x-y
}

func swap(a,b string) (string,string) {
	return b,a
}

// naked return
func split(sum int) (x,y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42,13))
	fmt.Println(substract(42,13))
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(17))
}