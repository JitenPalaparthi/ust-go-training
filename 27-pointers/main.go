package main

import (
	"errors"
	"fmt"
)

func main() {
	// shorthand declaration
	num := 100
	ptrnum := &num
	fmt.Println("Value of num:", num, "Value of num through pointer:", *ptrnum)
	fmt.Println("Address of num:", &num, "Address of num through pointer:", ptrnum)

	if err := changeVal(ptrnum, 200); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of num:", num, "Value of num through pointer:", *ptrnum)
	}
	if err := changeVal(&num, 300); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of num:", num, "Value of num through pointer:", *ptrnum)
	}

}

// using pointers as arguments is call by reference
func changeVal(variable *int, value int) error {
	if variable != nil {
		*variable = value
		return nil
	}
	return errors.New("cannot assign a value to a nil pointer")
}
