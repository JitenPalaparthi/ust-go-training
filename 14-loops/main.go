package main

import "fmt"

// for loop
// go does not have while or do-while loops
func main() {
	fmt.Println("1-10 numberss")
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println("Even numbers")
	for i := 2; i <= 10; i += 2 {
		fmt.Print(i, " ")
	}

	fmt.Println("\nIterating string")

	//str := "Hello, 世界!" // becasue some non english chars are not ascii chars; they are unicode chars
	str := "Hello, World!"
	//1- normal for loop , 2- for-range loop
	for i := 0; i < len(str); i++ {
		fmt.Print(string(str[i]), " ")
	}

	// reverse of string
	rstr := ""
	for i := 0; i < len(str); i++ {
		rstr = string(str[i]) + rstr
	}
	fmt.Println("\nreverse of str-->", rstr)
	// there is no char in go. rune is the type
}

// println , complex, len
// package, import, func , const, var , if , else, switch, for
