package main

import (
	"fmt"
	"unsafe"
)

// Byte-1 , Word-4, DWord-8 and QWord-16
// How does memory layout is formed?
// memory layout is created in the form of blocks.
// each block is 8 bytes
// if the whole struct can adjust in block then that memory is taken
// if it exceeds a block then it works as number of blocks (which are multiplied by 8)

func main() {
	fmt.Println("Size of T1 instance-->", unsafe.Sizeof(T1{}))
	// according to the layout it is 45 bytes
	fmt.Println("Size of T2 instance-->", unsafe.Sizeof(T2{}))
	fmt.Println("Size of T3 instance-->", unsafe.Sizeof(T3{}))

}

type T1 struct {
	//              	A    B     P  W
	bt1   byte   //	1    8     7
	char1 rune   //	2    7     5  5
	str1  string //	16  16     0
	bt2   byte   //	1    8     7  7
	num1  int64  //	8    8     0
	bool1 bool   // 	1    8     7  7
	str2  string // 	16  16     0
	//           	// 	45  64        19
	// A -> Actual
	// B-> Block and using Padding from the previous block
	// P -> Adjusted from Padding
	// W --> Wastage
}

type T2 struct {
	//              	A	B	P	W
	str2  string // 	16  16  0
	str1  string // 	16  16  0
	num1  int64  // 	8    8  0
	bt2   byte   // 	1    8  7
	bt3   byte   //	1    7  6
	bt1   byte   // 	1    6  5
	char1 rune   //	2    2  3   3
	//            		45   48
}

type T3 struct {
	i1   interface{}
	twob int16
	bt1  bool
	bt2  bool
	bt3  bool
	bt4  bool
	bt5  byte
	bt6  byte
}
