package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {

	fmt.Println(add(10, 20))
	fmt.Println(add(10.25, 20.25))
	fmt.Println(add("Hello", " World!"))
	fmt.Println(add(true, false))
	_, err := add(10, false)
	if err != nil {
		fmt.Println(err)
	}
}

// nil can be used with interface{}, pointer types, slice,map , chan

func add(a, b interface{}) (interface{}, error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return nil, errors.New("a and b are different types")
	}

	switch a.(type) {
	// case nil:
	// 	return nil, errors.New("nil type is not allowed")
	case int:
		return a.(int) + b.(int), nil
	case int8:
		return a.(int8) + b.(int8), nil
	case int16:
		return a.(int16) + b.(int16), nil
	case int32:
		return a.(int32) + b.(int32), nil
	case int64:
		return a.(int64) + b.(int64), nil
	case uint8:
		return a.(uint8) + b.(uint8), nil
	case uint16:
		return a.(uint16) + b.(uint16), nil
	case uint32:
		return a.(uint32) + b.(uint32), nil
	case uint64:
		return a.(uint64) + b.(uint64), nil
	case float32:
		return a.(float32) + b.(float32), nil
	case float64:
		return a.(float64) + b.(float64), nil
	case string:
		return a.(string) + b.(string), nil
	default:
		return nil, fmt.Errorf("add method cannot perform operation on the type of a:%T and b:%T", a, b)
	}

	return nil, nil
}
