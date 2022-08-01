package main

import "fmt"

//var val int
// defer func is called with argument, when defered the same value when called is used
func main() {

	val := 100
	// when you call the defer the val is 100.
	// So even defer is called the val 100 is maintained for defer call
	defer fmt.Println("Defer call:The Value of val:", val+1)
	val = val + 200
	fmt.Println("The Value of val:", val)

	str := "\nHello World!"
	for _, v := range str {
		defer fmt.Print(string(v))
	}

}

// func Calc() {
// 	val = val + 11
// }
