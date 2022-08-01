package main

import "fmt"

func main() {
	num := 15
	if num <= 50 {
		// func(n int) {
		// 	fmt.Println(n, "is less than or equal to 50")
		// 	return
		// }(num)
		fmt.Println(num, "is less than or equal to 50")
		return
	} else if num > 50 && num <= 100 {
		fmt.Println(num, "is less than 100 and greater than 50")
		return
	}

	fmt.Println(num, "is greater than 100")

	// if true {
	// 	fmt.Println("Its true")
	// } else {
	// 	fmt.Println("Its flase")
	// }
	// && || ! == !=
	//fmt.Println(num)

}
