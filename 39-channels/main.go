package main

import "fmt"

// 1- chan is a keyword to declare a channel
// 2- make is used to instantiate a channel
// 3- channels are buffered or unbuffered
// 4- in an unbuffered channel,
// 		- the sender is blocked until the receiver receives the value
// 		- the receiver is blocked until the sender sends the value
// 5- in a buffered channel,
//		- the sender is blocked if the channel buffer is full until at least one element of the buffer is freed
// 		- the receiver is also blocked if the channel buffer is full until one element of the buffer is freed
// 6- incase sender or receiver cannot send or receive then it is called dead lock.
// 		- the deadlock is either by sender or by receiver.
// 7- to send a value to the channel ch <- 100
// 8- to receive a value from the channel <- ch
func main() {
	var ch chan int
	ch = make(chan int) // unbuffered channel since size is not given
	// go func() {
	// 	ch <- 100
	// }() // sending a value from the channel
	// v := <-ch
	// fmt.Println("The value received from the channel", v)

	go func() {
		v := <-ch
		fmt.Println("The value received from the channel", v)
	}()
	ch <- 100
}
