package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)
	// go func() {
	// 	for i := 1; i <= 100; i++ {
	// 		time.Sleep(time.Millisecond * 10)
	// 		go publish("Publishing-->"+fmt.Sprint(i), ch)
	// 	}
	// }()

	// go func() {
	// 	for i := 1; i <= 100; i++ {
	// 		fmt.Println("Received-->", <-ch)
	// 	}
	// }()
	go publish(100, ch)

	// range loop on channels
	// for msg := range ch {
	// 	fmt.Println("Received-->", msg)
	// }
	go subscribe(ch)
	time.Sleep(time.Second * 3)
	fmt.Println("Finished receiving")
}

// func publish(message string, ch chan<- string) {
// 	ch <- message
// }

func publish(n int, ch chan<- string) {
	for i := 1; i <= n; i++ {
		time.Sleep(time.Millisecond * 10)
		ch <- "Publishing-->" + fmt.Sprint(i)
	}
	close(ch)
}

func subscribe(ch <-chan string) {
	for msg := range ch {
		fmt.Println("Received-->:", msg)
	}
}
