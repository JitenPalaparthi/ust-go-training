package main

import (
	"fmt"
	"math/rand"
)

// slice is a reference type
// when you create a slice it returns the pointer
// any reference type can check with nil
// the size of the slice can grow at runtime
// Other than shorthand declaration and instantiation, the usual way of instantiating
// 	slice by using make built in function
func main() {

	//arr1 := [3]int{10, 20, 30}  // this is an array
	slice1 := []int{10, 20, 30} // this is a slice; shorthand declaration
	fmt.Println(slice1, "Length", len(slice1), "Capacity", cap(slice1))

	var slice2 []int // This is only declaration not instantiation
	if slice2 == nil {
		fmt.Println("The slice2 yet to be instantiated")
	}
	//slice2 = make([]int, 2, 10)
	slice2 = make([]int, 10)
	fmt.Println(slice2, "Length", len(slice2), "Capacity", cap(slice2))
	slice2[0] = 10
	slice2[1] = 20
	fmt.Println(slice2, "Length", len(slice2), "Capacity", cap(slice2))
	for i := 0; i < len(slice2); i++ {
		slice2[i] = rand.Intn(1000)
	}
	fmt.Println(slice2, "Length", len(slice2), "Capacity", cap(slice2))
}
