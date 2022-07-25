package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// uintptr is an integer type that is large enough to hold the bit pattern of any pointer.
// an unsafe.Pointer can convert to uintptr.
// uintptr can be converted to unsafe.Pointer
// Why to convert to uintptr -->
// because arthemetic operation can be performed on an address that is stored in uintptr
func main() {
	arr := [6]int{10, 15, 18, 28, 34, 67}
	fmt.Println("Address of 0 , 1 and 2 elements in arr-->", &arr[0], &arr[1], &arr[2])
	ptr := uintptr(unsafe.Pointer(&arr[0]))
	size := unsafe.Sizeof(arr[0])
	for i := 1; i < len(arr); i++ {
		fmt.Println("Pointer:", ptr, "Type of pointer:", reflect.TypeOf(ptr), "Size of each element in bytes:", size)
		ptr = ptr + size // now ptr has the address of next element in the array
		val := (*int)(unsafe.Pointer(ptr))
		fmt.Println("Value at arr index ->", i, *val)
	}
}
