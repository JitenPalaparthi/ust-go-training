package main

import "fmt"

// empty interface{} --> can hold any kind of data
// similar to object in java/c#
// interface{} type can accept any kind of value
// type casting does not work between interface{} type to another type
// need to use type assertion which is v.(T)
// while doing type asserting , the underline type that interface{} has is important
func main() {

	var value interface{}

	value = 100 // assigning 100 to empty interface{}

	// var intvalue int = int(value) //Not Ok with empty interface T(v) --> Type casting
	var intvalue int = value.(int) // Ok with empty interface v.(T) --> Type assertion
	fmt.Println("interface{} Value", value, "Actual intvalue", intvalue)

	value = 100.001
	var float64value float64 = value.(float64)
	fmt.Println("interface{} Value", value, "Actual float64value", float64value)

	// caste(assertion) ==> T(v.(T))
	var float32value float32 = float32(value.(float64))
	fmt.Println("interface{} Value", value, "Actual float32value", float32value)

	value = true
	var boolvalue bool = value.(bool)
	fmt.Println("interface{} Value", value, "Actual boolvalue", boolvalue)

	value = "Hello World!"
	var stringvalue string = value.(string)
	fmt.Println("interface{} Value", value, "Actual stringvalue", stringvalue)

}
