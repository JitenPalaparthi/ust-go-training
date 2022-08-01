package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello World")
	}()

	func(a, b int) {
		fmt.Println("Addition of a and b", a+b)
	}(10, 20)

	f := func(a int, b int) {
		fmt.Println("Addition of a and b", a+b)
	}

	f(10, 20)

	k := func(a, b int) int {
		return a + b
	}(10, 20)

	fmt.Println("Addition:", k)
}
