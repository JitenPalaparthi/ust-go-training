package main

import "fmt"

// buffered channel does not block the caller or callee until the buffer is full
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch, <-ch)
	ch <- 3
	fmt.Println(<-ch)
}
