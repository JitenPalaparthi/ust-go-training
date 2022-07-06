package main

import (
	"fmt"
)

func main() {
	var arr1 [10]int
	arr2 := [3]int{10, 12, 13}           // short hand declaration of array
	arr3 := [...]int{10, 11, 12, 13, 14} // declaring an array without mentioning length
	//var arr4 [...]int // Not OK.
	fmt.Println(arr1, arr2, arr3)

	// for range loop
	//foreach(int a in arr2)
	//for(int a:arr2)
	// for-range loop can be used in arrays, slices , maps and channels
	// the behaviour of range loop is different for maps and channels
	// for maps it returns key and value
	// for channels it returns only one value
	// for arrays and slices it returns index and value
	for i, v := range arr3 {
		fmt.Println("Index", i, "Value", v)
	}
	str := "Hello, 世界"

	// the normal for loop iterates as a byte
	fmt.Println("--- Using for loop on string iteration-----")
	for i := 0; i < len(str); i++ {
		fmt.Print("\nByte-->", str[i], "Char-->", string(str[i]))
	}

	fmt.Println("\n-- Using range loop on string iteration----")

	for _, v := range str {
		fmt.Print("\nRune-->", v, "Char-->", string(v))
	}

}
