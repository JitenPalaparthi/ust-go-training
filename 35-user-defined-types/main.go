package main

import "fmt"

func main() {

	var mp1 myMap // delcare a varialbe called mp1 of type myMap

	mp1 = make(map[string]interface{})

	mp1["560086"] = "Yeshvantpur, Bangalore"
	mp1["560096"] = "Mahalakshmi Layout, Bangalore"
	mp1["522001"] = "Guntur-1"
	mp1["3423423"] = "Toronto, Canada"
	mp1["332243"] = "Chicago, Usa"

	fmt.Println(mp1.GetKeys())

	var mp2 map[string]interface{}
	mp2 = mp1
	fmt.Println(mp2)
	fmt.Println("Keys in a map", myMap(mp2).GetKeys())
}

type myMap map[string]interface{}

func (m myMap) GetKeys() []string {
	keys := make([]string, 0)
	for key, _ := range m {
		keys = append(keys, key)
	}
	return keys
}
