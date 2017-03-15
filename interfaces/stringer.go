package main

import "fmt"

/*
 * type Stringer interface {
 *    String() string
 * }
 */

type Person struct {
	Name string
	Age  int
}

// pointer receiver won't be called by default!
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years old)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a,z)
}