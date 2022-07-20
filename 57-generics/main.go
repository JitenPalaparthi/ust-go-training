package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Addition:", additionNonGeneric(12, 15))
	fmt.Println("Addition:", additionNonGeneric(12.45, 15.36))
	fmt.Println("Addition using Generics:", additionGeneric(12.45, 15.36, 12.345))
	fmt.Println("Addition using Generics:", additionGeneric(12345, 4567, 232))
}

type NType interface {
	int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | int | float32 | float64
}

//func additionGeneric[N int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | int | float32 | float64 | string](nums ...N)N{
func additionGeneric[N NType](nums ...N) N {
	var sum N
	for _, v := range nums {
		sum += v
	}
	return sum
}

func additionNonGeneric(a, b interface{}) interface{} {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return nil
	}
	switch a.(type) {
	case int:
		return a.(int) + b.(int)
	case float32:
		return a.(float32) + b.(float32)
	}
	return nil
}
