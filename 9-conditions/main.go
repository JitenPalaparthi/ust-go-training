package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// if else
	var age uint8 = 17
	if age >= 18 {
		fmt.Println("Eligible for vote")
	} else {
		fmt.Println("Not eligible for vote")
	}

	//num := 25

	if num := rand.Intn(51); num%2 == 0 {
		fmt.Println("The number", num, "is even")
	} else {
		fmt.Println("The number", num, "is odd")
	}
}

// expressions vs statements
