package main

import (
	"errors"
	"fmt"
)

func main() {
	slice1 := make([]int, 2)
	// slice1[0]=0
	// slice1[1]=0
	fmt.Println(slice1, "Len", len(slice1), "Cap", cap(slice1))
	slice1[0] = 100
	//[0,0,0,0,0]
	fmt.Printf("Address of slice1 %p\n", &slice1)
	slice1 = append(slice1, 1)
	slice1 = append(slice1, 2)
	slice1 = append(slice1, 3)
	slice1 = append(slice1, 4)
	slice1 = append(slice1, 5)
	slice1 = append(slice1, 6)
	//slice1[0]=1
	//[0,0,0,0,0,1]
	fmt.Println(slice1)
	fmt.Printf("Address of slice1 %p\n", &slice1)

	var slice2 []int           // only declared slice
	slice2 = append(slice2, 1) // is the slice is nil, then the append instantiates and appends a value
	slice2 = append(slice2, 2)
	slice2 = append(slice2, 3)
	slice2 = append(slice2, 4)
	slice2 = append(slice2, 5)
	fmt.Println(slice2, "Len", len(slice2), "Cap", cap(slice2))

	slice3 := slice2 // slice3 and slice2 refer to the same memory location

	slice3[0] = 100
	slice3[1] = 101
	slice3[2] = 101
	if err := changeVal(slice2, 4, 501); err != nil {
		fmt.Println(err)
	}
	slice2 = append(slice2, 600)
	fmt.Println("Slice2", slice2, "Len", len(slice2), "Cap", cap(slice2))
	fmt.Println("Slice3", slice3, "Len", len(slice3), "Cap", cap(slice3))

	var slice4 []int
	slice4 = make([]int, len(slice2))
	copy(slice4, slice2) // to copy a slice first the destination slice must not be nil
	// for i := 0; i < len(slice2); i++ {
	// 	slice4[i] = slice2[i]
	// }
	fmt.Println("Slice4", slice4, "Len", len(slice4), "Cap", cap(slice4))
	changeVal(slice4, 0, 1100)
	fmt.Println("Slice2", slice2, "Len", len(slice2), "Cap", cap(slice2))
	fmt.Println("Slice4", slice4, "Len", len(slice4), "Cap", cap(slice4))
}

func changeVal(slice []int, index, val int) error {
	if index >= len(slice) {
		errors.New("out of bounds")
	}
	slice[index] = val
	return nil
}
