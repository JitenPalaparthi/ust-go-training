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
	duration := time.Now().Add(time.Second * 1)
	rootCtx := context.Background()
	ctx, cancel := context.WithDeadline(rootCtx, duration)
	defer cancel()
	go square(ctx, ch)
	//go square(ctx, ch)

	fmt.Println("Number of CPUs,goroutines running", runtime.NumCPU(), runtime.NumGoroutine())

	for v := range ch {
		fmt.Println("The square of i->", v)
	}

	//time.Sleep(time.Second * 1)
	fmt.Println("Number of CPUs,goroutines running", runtime.NumCPU(), runtime.NumGoroutine())
}

func square(ctx context.Context, out chan int) {
	i := 0
	for {
		//time.Sleep(time.Second * 1)
		select {
		case <-ctx.Done():
			close(out)
			return
		case out <- i * i:
			i++
		}
	}
}
