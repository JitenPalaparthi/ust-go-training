package main

import (
	"fmt"
	"time"
)

// 1.to create a go routine just need to use a keyword go infront of any method/function
// 2.main is also a go routine
// 3.by nature, no go rounrines waits for other go routine(s) to complete it's execution
// 4.by default, the order or exececution cannot be determined
func main() {
	go greet()

	go fmt.Println("Hello World!", " This is called in main")

	go func(i, j int) {
		fmt.Println("Addition:", i+j)
	}(100, 200)
	//var k int
	go func() {
		k := add(10, 20)
		fmt.Println("Calling add func:", k)
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			go fmt.Println(" I:", i)
		}
	}()

	time.Sleep(time.Millisecond * 1)

}

func greet() {
	fmt.Println("Hello World!", "This is called as a go routine")
}

func add(i, j int) int {
	return i + j
}
