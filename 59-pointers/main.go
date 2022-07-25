// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var num = 100
	var iptr uintptr = 0x0000ffff

	arr := [3]int{10, 13, 14}

	fmt.Println(&arr[0], &arr[1], &arr[2])

	iptr = uintptr(unsafe.Pointer(&arr[0]))
	size := unsafe.Sizeof(arr[0])
	fmt.Println(iptr)
	iptr = iptr + size
	ptr := (*int)(unsafe.Pointer(iptr))

	fmt.Println(num, iptr, *ptr)
	e1 := Employee{name: "Jitendranadh Palaparthi", email: "jitenp@outlook.com"}
	fmt.Println(unsafe.Sizeof(e1))
	fmt.Println(unsafe.Sizeof(e1.name))
	fmt.Println(unsafe.Sizeof(e1.email))
	fmt.Println(unsafe.Sizeof(e1.id))

	//var str string = "Hello This is Jiten, I am trying to understand how does this work"
	fmt.Println("-->", unsafe.Alignof(e1))
	fmt.Println("-->", unsafe.Offsetof(e1.contactno))
	fmt.Printf("required alignment: %+v\n", unsafe.Alignof(Foo{})) // required alignment: 4
	fmt.Printf("required alignment: %+v\n", unsafe.Alignof(Bar{})) // required alignment: 8

}

type Employee struct {
	id        bool
	ok        bool
	contactno int
	address   string
	name      string
	email     string
}
type Foo struct {
	aaa [2]bool
	bbb int32
	ccc [2]bool
}

type Bar struct {
	aaa [2]bool
	ccc [2]bool
	bbb int64
}
