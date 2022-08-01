package main

import "fmt"

// to declare a map  var mp map[string]string. map[kT]VT KT is key Type and VT is value type
// use make to instantiate the map
// What types can be key type?  Which ever the type that can perform == operation can be a key type

func main() {

	// var num int
	// const c int = 10
	// var arr [3]int
	// var slice []int

	var mp map[string]string
	// mp = make(map[string]string, 10) OK
	// mp := make(map[string]string) shorthand declaration
	if mp == nil {
		mp = make(map[string]string)
	}

	//var mapbool map[bool]string OK
	//var mapfloat map[float32]string OK
	// var imap map[interface{}]interface{}
	// imap = make(map[interface{}]interface{})
	// imap["name"] = "Jiten"
	// imap[37] = "age"
	// imap[true] = "male"
	// fmt.Println(imap)
	mp["560086"] = "Yeshvantpur , Bangalore"
	mp["560096"] = "Mahalakshmi Layout, Bangalore"
	mp["522001"] = "Guntur-1"
	mp["690011"] = "Trivandrum Medical College"

	fmt.Println(mp)

	// iterate through the map

	fmt.Println("Length of map", len(mp))
	//fmt.Println("Capacity of map", cap(mp)) cap cannot work on map

	// iteration of map
	// range loop returns key and value in map not the index and value
	for key, value := range mp {
		fmt.Println("Key", key, "Value", value)
	}
	// delate a key from a map
	fmt.Println("Delete key 522001")
	delete(mp, "522001")
	for key, value := range mp {
		fmt.Println("Key", key, "Value", value)
	}
	//
	val, ok := mp["522001"] // map returns value and key exists or not
	fmt.Println("Is key exists", ok, "Value of a key", val)
	val, ok = mp["560086"] // map returns value and key exists or not
	fmt.Println("Is key exists", ok, "Value of a key", val)

	// a nil can be checked on slice, map, chan , interface{} and any pointertype

}
