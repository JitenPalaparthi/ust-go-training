package main

import (
	"time"
)

func main() {

	ch := make(chan int)
	//done := make(chan struct{})
	//go func() { ch <- 1 }()
	// go func() {
	// 	for i := 1; ; i++ {
	// 		go func() {
	// 			v := <-ch
	// 			ch <- v + 1
	// 			println("-->", v)
	// 			if v == 999 {
	// 				fmt.Println("Done???")
	// 				done <- struct{}{}
	// 			}
	// 		}()
	// 	}
	// }()
	go func() {
		//ch <- 0
	}()
	for i := 1; i <= 1000; i++ {
		go incr(ch)
	}
	time.Sleep(time.Second * 2)
	//<-done
	//fmt.Println("Done")
}

func incr(ch chan int) {
	//var v int
	go func() {
		v := <-ch
		println(v)
	}()

	//v := <-ch

}
