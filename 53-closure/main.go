package main

import "fmt"

func main() {
	fmt.Println("Addition", calc(10, 20, add))
	fmt.Println("Substract", calc(10, 20, sub))
	fmt.Println("Multiply", calc(10, 20, mult))
	fmt.Println("Division", calc(20, 10, func(a, b int) int {
		return a / b
	}))
}
func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}
func calc(a, b int, f func(int, int) int) int {
	fmt.Println("Performaing caluclation-->")
	return f(a, b)
}
