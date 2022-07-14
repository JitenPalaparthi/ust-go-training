package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	message := make(chan string)
	go greet(done)
	//go sender(message, "Hello World")
	//go receiver(message)
	for i := 1; i <= 100; i++ {
		go sender(message, "Hello World--> "+fmt.Sprint(i))
	}
	for i := 1; i <= 100; i++ {
		go receiver(message)
	}
	<-done
	fmt.Println("Hello world from main function")
}

func greet(done chan bool) {
	fmt.Println("Hello World from greet function")
	time.Sleep(time.Millisecond * 1)
	done <- true
}

func sender(send chan<- string, message string) { // send only channel. i.e can send data but cannot receive data
	send <- message
	// <-send this operation is forbidden
}

func receiver(receive <-chan string) { // receive only channel.i.e cannot send data but can receive
	fmt.Println(<-receive)
	// receive <- "Hello world" this operation is forbidden
}
