package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var num1 int = 100

	str1 := fmt.Sprint(num1)
	fmt.Println("Value and type of Str", str1, reflect.TypeOf(str1))

	// 2- using strconv

	str2 := strconv.Itoa(num1)
	fmt.Println("Value and type of Str", str2, reflect.TypeOf(str2))

	var num2 myInt = 100

	fmt.Println("Value and type of num2 methods;ToString and ToByte", num2.ToString(), reflect.TypeOf(num2.ToString()), num2.ToByte(), reflect.TypeOf(num2.ToByte()))

	var num3 int = 200

	fmt.Println("Covert from int to myInt and call ToString:", myInt(num3).ToString())

}

type myInt int

// 1- You can implement user defined types from the existing types.
//	So can create methods on them as well.

// 2- It is very useful when interfaces to be passed as parameters.To acheive DI,IOC or similar patterns

func (mi myInt) ToString() string {
	return fmt.Sprint(mi)
}

func (mi myInt) ToByte() []byte {
	return []byte(fmt.Sprint(mi))
}
