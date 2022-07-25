package main

import (
	_ "fmt"
)

//var i uint8 = 10
const (
	MIN = 1
	// MAX1, MIN1 = 10 * i, 1 // Not OK this is wrong becasuse i is a normal variable
	MAX1, MIN1 = 10 * MIN, 1 // OK
	MIX2       = 10 * 10     // OK
)

func main() {
	const MAX int = 100

	println("MAX1:", MAX1)
	// _ blank identifier
}
