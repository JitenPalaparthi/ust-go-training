package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	var ch = make(chan int)
	//ctx := context.Background() // creating a root channel
	ctx, cancel := context.WithCancel(context.Background())
	go square(ctx, ch)

	fmt.Println("Number of CPUs,goroutines running", runtime.NumCPU(), runtime.NumGoroutine())

	for i := 0; i < 10; i++ {
		fmt.Println("The square of i->", <-ch)
	}
	cancel()

	time.Sleep(time.Second * 1)
	fmt.Println("Number of CPUs,goroutines running", runtime.NumCPU(), runtime.NumGoroutine())
}

func square(ctx context.Context, out chan int) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			return
		case out <- i * i:
			i++
		}
	}
}
