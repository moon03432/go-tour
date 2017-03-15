package main

import "fmt"

/*
 * map: key-value collection
 */
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

// direct assignment
var m2 = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {

	fmt.Println(m2)

	// make: create maps
	m := make(map[string]int)

	// add, update, delete, get
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)


}
