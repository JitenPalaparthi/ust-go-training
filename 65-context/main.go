package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	//var ch = make(chan int)
	deadline := time.Now().Add(time.Second * 2)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	go worker(ctx, 3)
	go worker(ctx, 1)
	go worker(ctx, 2)

	fmt.Println("Number of CPUs,goroutines running", runtime.NumCPU(), runtime.NumGoroutine())

	time.Sleep(time.Second * 4)
	fmt.Println("Number of CPUs,goroutines running", runtime.NumCPU(), runtime.NumGoroutine())
}

func worker(ctx context.Context, seconds int) {
	select {
	case <-ctx.Done():
		fmt.Println("Killed the worker")
		return
	case <-time.After(time.Second * time.Duration(seconds)):
		fmt.Println("Work done--->")
	}
}
