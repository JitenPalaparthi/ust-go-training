package main

import (
	"fmt"
	"reflect"
)

func main() {

	c1 := complex(12.34, 10.45)
	// two variations of complex numbers 64 and 128
	fmt.Println("Complex number-1", c1, " Type", reflect.TypeOf(c1))

	var r1, i1 float32 = 12.34, 10.45
	c2 := complex(r1, i1)
	fmt.Println("Complex number-2", c2, " Type", reflect.TypeOf(c2))

	c3 := 12.34 + 10.45i
	fmt.Println("Complex number-3", c3, " Type", reflect.TypeOf(c3))

	c4 := c1 + c3

	fmt.Println("Addition of complex numbers:", c4)

	c5 := c1 - c3

	fmt.Println("Substraction of complex numbers:", c5)

	c6 := c1 * c3

	fmt.Println("Multiplication of complex numbers:", c6)

	c7 := c1 / c3

	fmt.Println("Division of complex numbers:", c7)

	// c8 := c1 + c2 // This is not possible due to mismatched types

}
