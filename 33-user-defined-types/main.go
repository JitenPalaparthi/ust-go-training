package main

import (
	"fmt"
)

func main() {

	// One of the ways to achieve polymorphysm
	g1 := Greeting{}

	g1.Greet()

	g2 := Greeting1{}

	g2.Greet("Hello World")
}

type Greeting struct{} // empty structure

// A method for an empty structure
func (e Greeting) Greet() {
	fmt.Println("Hello World")
}

func Greet(str string) {
	fmt.Println(str)
}

type Greeting1 struct{}

func (g Greeting1) Greet(str string) string {
	return str
}
