package main

import (
	"fmt"
	"math/rand"
)

func main() {

	slice1 := make([]int, 10)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = rand.Intn(100)
	}

	fmt.Println(slice1)

	slice2 := slice1[:5] // from 0(beginning) to 5th
	fmt.Println(slice2)

	slice3 := slice1[3:7] // from 3rd to 7th
	fmt.Println(slice3)

	slice4 := slice1[5:]
	fmt.Println(slice4) // from 5th to the end

	slice4[0] = 1111
	fmt.Println(slice4)
	fmt.Println(slice1)
	//slice1 = append(slice1, 1, 2, 100, 101, 232)
	slice5 := remove(slice1, 4)
	fmt.Println(slice5)
}

func remove(slice []int, index int) []int {
	if index < len(slice) {
		slice2 := slice[:index-1]
		slice3 := slice[index:]
		for i := 0; i < len(slice3); i++ {
			slice2 = append(slice2, slice3[i])
		}
		return slice2
	}
	return nil
}

// strings
// sort
