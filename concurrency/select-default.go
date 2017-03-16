package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100*time.Millisecond)
	boom := time.After(500*time.Millisecond)

	for {
		select {
		case <- tick:
			fmt.Print("tick.")
		case <- boom:
			fmt.Print("BOOM!")
			return
		default:
			fmt.Print("    .")
			time.Sleep(50*time.Millisecond)
		}
	}
}