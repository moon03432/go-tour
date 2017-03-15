package main

import "fmt"

type Vertex struct {
	X,Y float64
}

func main()  {
	// constructor
	v := Vertex{1,2}
	fmt.Println(v)

	// accessors
	v.X = 4
	fmt.Println(v)

	// pointer
	p:= &v
	p.X = 1
	fmt.Println(v)

	// implicit values
	v1 := Vertex{X:1}
	v2 := Vertex{}
	p = &Vertex{1,2}
	fmt.Println(v1,v2,p)	// {1 0} {0 0} &{1 2}
}