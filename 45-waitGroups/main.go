package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go greetBy(&wg)
	fmt.Println("Hello world from main")
	wg1 := sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg1.Add(1)
		go func(w *sync.WaitGroup) {
			greet()
			w.Done()
		}(&wg1)
	}
	wg.Wait()
	wg1.Wait()

}

func greetBy(wg *sync.WaitGroup) {
	fmt.Println("Hello World! from greetBy WaitGroup func")
	wg.Done()
}

func greet() {
	fmt.Println("Hello World! from greet func")
}
