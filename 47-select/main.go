package main

import (
	"fmt"
	"time"
)

func main() {
	select {
	case v1 := <-timeOut():
		fmt.Println("It is timed out", v1)
	case v2 := <-retry():
		fmt.Println("retried", v2)
	}
	time.Sleep(time.Second * 5)
}

func timeOut() chan bool {
	ch := make(chan bool)
	go func() {
		fmt.Println("checking for timedout")
		time.Sleep(time.Second * 2)
		ch <- true
	}()
	return ch
}

func retry() chan bool {
	ch := make(chan bool)
	go func() {
		fmt.Println("retrying to connect")
		time.Sleep(time.Second * 3)
		ch <- true
	}()
	return ch
}
