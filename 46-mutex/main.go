package main

import (
	"sync"
)

var counter *int

func main() {
	//var counter int
	n := 0
	counter = &n
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(2)
	go func() {
		for i := 1; i <= 2000; i++ {
			wg.Add(1)
			go incr(wg, mu)
		}
		wg.Done()
	}()

	go func() {
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			go incr(wg, mu)
		}
		wg.Done()
	}()
	wg.Wait()
	println(n)
}

func incr(wg *sync.WaitGroup, mu *sync.Mutex) {
	//counter := 0
	mu.Lock()
	*counter++
	mu.Unlock()
	wg.Done()
}
