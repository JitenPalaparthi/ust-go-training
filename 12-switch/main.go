package main

import "fmt"

func main() {
	// Which ever scenarios , the break is removed in java/c# in go you need to use fallthrough

	//num := 16
	switch num := 16; {
	case num%8 == 0:
		fmt.Println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		fmt.Println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		fmt.Println(num, "is divisible by 2")
	default:
		fmt.Println("not divisible by 8 4 or 2")
	}
}
