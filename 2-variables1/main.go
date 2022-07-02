package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var num1 int
	println("Number1:", num1)
	num1 = 1
	println("Number1:", num1)
	var num2 int = 10
	println("Number2:", num2)

	var num3 = 20 // 8 bytes
	println("Number3:", num3)
	// num3 = "10"
	// Q: What is the size of int? 4 or 8 32bit arch --> 4bytes 64bit arch-->8bytes
	fmt.Println("Size of Int", unsafe.Sizeof(num3))
	fmt.Println("Type of num3", reflect.TypeOf(num3))

	// int8, int16, int32, int64,uint8,uint16,uint32,uint64

	// float32 , float64

	//var float1 float32

	var float1 = 12.45
	fmt.Println("Type of float1", reflect.TypeOf(float1))
	println(float1)
	var float2 float64
	println(float2)

	var str string = "Hello World"

	println(str)

	var (
		a, b, c int = 10, 20, 30
	)

	fmt.Printf("A:%d B:%d C:%d", a, b, c)

	//declaring and assigning values to group of variables
	var (
		p, q, r = 10, 10.10, false
	)

	fmt.Printf("P:%d Q:%f R:%v\n", p, q, r)

	str1 := "Hello World" // shorthand declaration

	fmt.Printf("type of %T", str1)

	var (
		a1, b1 = 10, 20
	)
	fmt.Println("A1:", a1, "B1:", b1)
	// general swaping
	// t := a1
	// a1 = b1
	// b1 = t
	a1, b1 = b1, a1 // swap
	fmt.Println("A1:", a1, "B1:", b1)

	var rn rune = 95
	fmt.Printf("Rune %v\n", rn)
	fmt.Println("Hello, 世界")
	var bt byte = 96
	fmt.Printf("Byte %v\n", bt)
}
