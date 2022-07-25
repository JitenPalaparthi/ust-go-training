package main

import (
	"encoding/json"
	"fmt"
)

// move to heap: The data is stored in heap memory
// escape to heap: particular variable has to be shared between main and other function or method
// does not escape to heap: there is a less burden on GC
func main() {
	val := getVar()
	ptr1 := getPtr()
	ptr2 := &val
	changeVal(ptr2)
	//fmt.Println(ptr1, val)
	println(ptr1, val)
	e1 := &Employee{Name: "Jiten", Email: "JitenP@Outlook.Com"}
	jsonbuf, err := json.Marshal(e1)
	println("marshal from struct to json")
	if err != nil {
		println(err)
	} else {
		println(string(jsonbuf))
	}
	println("unmarshal from json to struct")
	var e2 *Employee
	err = json.Unmarshal(jsonbuf, e2)
	if err != nil {
		println(err)
	} else {
		fmt.Println(e2)
	}
}

func getPtr() *int {
	// num := new(int)
	// *num = 100
	// return num
	num := 100
	return &num
}

func getVar() int {
	num := 100
	return num
}

func changeVal(ptr *int) {
	*ptr = 200
}

type Employee struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
