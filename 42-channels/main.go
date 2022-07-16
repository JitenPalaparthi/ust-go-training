package main

import (
	"fmt"
)

func main() {

	val := <-notifyMe()
	fmt.Println(val)
	fmt.Println("Called in main")
}

func notifyMe() chan bool {
	done := make(chan bool)
	go func() {
		//time.Sleep(time.Second * 5)
		for i := 1; i <= 100; i++ {
			fmt.Print(i, " ")
		}
		done <- true
	}()
	return done
}
