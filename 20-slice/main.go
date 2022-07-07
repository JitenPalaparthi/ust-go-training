package main

import "fmt"

func main() {
	slice1 := make([]int, 6, 6)
	slice1[0] = 10
	slice1[1] = 11
	slice1[2] = 12
	slice1[3] = 13
	slice1[4] = 14
	slice1[5] = 15
	fmt.Println("Slice1", slice1, "length", len(slice1), "Capacity", cap(slice1))
	slice1 = append(slice1, 16)
	fmt.Println("Slice1", slice1, "length", len(slice1), "Capacity", cap(slice1))
	slice1 = append(slice1, 17)
	slice1 = append(slice1, 18)
	slice1 = append(slice1, 19)
	slice1 = append(slice1, 20)
	slice1 = append(slice1, 21)
	slice1 = append(slice1, 22)
	fmt.Println("Slice1", slice1, "length", len(slice1), "Capacity", cap(slice1))
	slice1 = append(slice1, 23)
	fmt.Println("Slice1", slice1, "length", len(slice1), "Capacity", cap(slice1))
}
