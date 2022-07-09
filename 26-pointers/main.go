package main

import "fmt"

func main() {
	var num int = 100
	var ptrnum *int

	// can check a pointer is nil or not

	if ptrnum == nil {
		// When the pointer is nil , it cannot reference any value becasue it does not hold any address
		// fmt.Println("Value that pointer holds:", *ptrnum)
		// nil pointer dereference: When the pointer does not points to any address
		ptrnum = &num
	}

	fmt.Println("Value of num:", num, "Value of num through pointer:", *ptrnum)
	fmt.Println("Address of num:", &num, "Address of num through pointer:", ptrnum)

	*ptrnum = 200 // This reassigns the value inside the variable it points to
	fmt.Println("Value of num:", num, "Value of num through pointer:", *ptrnum)

	ptrnum = nil // The pointer is not points to any address
	fmt.Println("ptrnum after assigning nil:", ptrnum)
}
