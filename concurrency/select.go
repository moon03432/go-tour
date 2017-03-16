package main

import (
	"fmt"
	"time"
)

/*
 * The select statement lets a goroutine wait on multiple communication operations.
 *
 * A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
 */
func fibonacci(c,quit chan int) {
	x,y := 0,1

	for {
		select {
		case <- quit:
			fmt.Println("quit")
			return
		case c <- x:
			x,y = y, x+y
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i<10 ; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	time.Sleep(time.Second)

	fibonacci(c,quit)
}