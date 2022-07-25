package main

import (
	"errors"
	"fmt"
)

func main() {

	// usually we need a variable that points to a pointer
	// but new function creates a pointer and assign a default value to that pointer
	ptr := new(int)
	fmt.Println("Address and value", ptr, *ptr)
	changeVal(ptr, 100)
	fmt.Println("Address and value", ptr, *ptr)

	// pointer of pointer. This holds the addres of another pointer using **  notation

	var ptrOfptr **int = &ptr
	fmt.Println("Address and value of ptr", ptr, *ptr)
	fmt.Println("Address and value of ptrOfptr", *ptrOfptr, **ptrOfptr)

	changeVal(*ptrOfptr, 300)
	fmt.Println("Address and value", ptr, *ptr)

	var ptrOfptrOfptr ***int = &ptrOfptr
	fmt.Println("Address and value of ptr", ptr, *ptr)
	fmt.Println("Address and value of ptrOfptr", *ptrOfptr, **ptrOfptr)
	fmt.Println("Address and value of ptrOfptrOfptr", **ptrOfptrOfptr, ***ptrOfptrOfptr)
	// ptr++ Not Ok. On pointers by default cannot perform arithmetic operations on pointers
}

func changeVal(ptr *int, val int) error {
	if ptr != nil {
		*ptr = val
		return nil
	}
	return errors.New("cannot assign a value to a nil pointer")
}
