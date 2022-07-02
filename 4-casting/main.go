package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// There is not implicit type casting in Go
	// int and int64 are not same even both carries 8 bytes on 64bit arch machines
	// int8 and uint8 are not same. Must be type casted
	var age1 uint8 = 37      // 1 byte
	var age2 int = int(age1) // not OK without type casting 8 bytes
	// type casting T(v) not X(T)v
	// age2 := age
	var age3 int8 = int8(age1) // Not Ok unlesss type casted
	// twitter
	// 1 million active users on twitter
	// age instead of uint8 created as int --> 7 bytes is wasted. 7 million bytes
	fmt.Println("Age1:", age1, "Age2:", age2, "Age3", age3)

	var float1 float32 = 10.45
	var num1 int = int(float1) // this does not change the type of a variable. It only accomidate the value to fit in
	//float1 = num1              // Not OK
	float1 = float32(num1) //Ok

	fmt.Println("Float1:", float1, "Num1:", num1)

	var num2 uint8

	num3 := 1234 /// type int

	num2 = uint8(num3) // 1234 is bigger then the range of int8 still it tries to fit but can not properly
	fmt.Println("Num2", num2, "Num3", num3)

	// string conversions

	runetype := 'A' // This is taken as rune
	var num5 int = int(runetype)

	fmt.Println("rune", runetype, "Num5", num5)

	num4 := 65
	var str string = string(num4)
	// I want to conver 65 as "65" not as unicode A
	// What is the result?
	fmt.Println("Str", str, "Num4", num4)

	str2 := fmt.Sprint(num4) // Sprint converts any thing to string

	fmt.Println("String conversion and type", str2, reflect.TypeOf(str2))

	str3 := fmt.Sprint(true) // Sprint converts any thing to string

	fmt.Println("String conversion and type", str3, reflect.TypeOf(str3))

	str4 := "66"

	num6, _ := strconv.Atoi(str4)

	fmt.Println("Num6 and type", num6, reflect.TypeOf(num6))

	num7 := 889
	fmt.Println("Integer to String:", strconv.Itoa(num7))

}
