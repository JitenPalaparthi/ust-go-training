package main

import "fmt"

func main() {

	// map of array as value
	// var mymap map[string][3]string

	// mymap = make(map[string][3]string)
	// mymap["Bangalore"] = [3]string{"560086", "560096", "560031"}
	// mymap["Guntur"] = [3]string{"522001", "522002", "522003"}

	// map slice as a value
	var mymap map[string][]string
	mymap = make(map[string][]string)
	mymap["Bangalore"] = []string{"560086", "560096", "560031"}
	mymap["Guntur"] = []string{"522001", "522002", "522003"}
	mymap["Hyderabad"] = nil
	val, ok := mymap["Hyderabad"]

	// if the value can be a type can be nil then can check
	if val == nil && ok {
		fmt.Println("val is nil but key exits.Val and is key exists", val, ok)
	}

}
