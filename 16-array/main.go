package main

import (
	"fmt"
	"reflect"
)

// Arrays in go are value types
// The type of array contains its length as well.
// Arrays are immutable. The length of the array must be known at complile time

func main() {
	var arr [3]int

	arr[0] = 100
	arr[1] = 101
	arr[2] = 102
	fmt.Println("Array and type of an array", arr, reflect.TypeOf(arr))
	var arr2 [4]int
	arr2[0] = 100
	arr2[1] = 101
	arr2[2] = 102
	fmt.Println("Array and type of an array", arr2, reflect.TypeOf(arr2))
	//sumOfArr(arr2) // This is not possible because arr2 is not same type as the function parameter

	for i := 0; i < len(arr); i++ {
		fmt.Println("Array Index", i, "Value", arr[i])
	}

}

func sumOfArr(arr [3]int) int {
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
