package main

import "fmt"

func main() {
	//fmt.Println("Hello", "World!", "How", "are", "you", "doing?")
	// fmt.Println("Hello World! How are you doing?")

	fmt.Println("SumOf:", sumOf(10, 20, 30))
	fmt.Println("SumOf:", sumOf(10, 20, 30, 40, -10))
	fmt.Println("SumOf:", sumOf(10, 20, 30, 40, -10, -20, -30, -40))
	fmt.Println("SumOf:", sumOf())

	slice := []int{10, 20, 30, 40, 50}
	fmt.Println("SumOf slice:", sumOf(slice...))

	arr := [3]int{-10, -20, -30}
	// sliceArr := arr[0:len(arr)] // convert an array to slice
	// sliceArr := arr[:] // the above one and this are same
	fmt.Println("SumOf arr:", sumOf(arr[:]...)) //arr[:] converts an array to slice
	// arr[:] conversion can be used as and where slice is the parameter

	//sumOfAny(10, 10.5, "Hello", true, 1, 0, -10.34)

}

// eclipse symbol ... is used to create a variadic parameter
// the variadic parameter must be the last parameter. So only one variadic parameter can be used
// len can used to determine the number of elements
// zero arguments can be passed as well
// a slice can be an argument to the variabic parameter. Use slice... provided by slice is variable name of type []T

func sumOf(nums ...int) int {
	sum := 0
	// for i := 0; i < len(nums); i++ {
	// 	sum = sum + nums[i]
	// }
	for _, v := range nums {
		sum += v
	}
	return sum
}

// Stub for the task
func sumOfAny(vals ...interface{}) interface{} {

	return nil
}
