package main

import "fmt"

func main() {
	// buffered channel has size. If buffer is full, there is a fatal error (deadlock)
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println(<-ch, <-ch)
}