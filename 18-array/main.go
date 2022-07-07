package main

import (
	"errors"
	"fmt"
)

func main() {

	arr1 := [3]int{100, 101, 102}

	fmt.Println(arr1)

	if err := changeVal(arr1, 3, 201); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Array after calling changeValue func", arr1)
	}

	//fmt.Println(arr1[3])

	/*for i := 0; i <= len(arr1); i++ {
		fmt.Println(arr1[i])
	}*/
	// panic is an error that is not usually recovered.
	// panic crashes your application

	// 2D Array

	arr2 := [3][2][2]int{{{1, 2}, {11, 22}}, {{3, 4}, {33, 44}}, {{5, 6}, {55, 66}}}
	fmt.Println("3D array", arr2, "Length of arr2", len(arr2))

	// This prints only 2D of 3D array
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(arr2[i][j])
		}
	}

	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			for k := 0; k < len(arr2[i][j]); k++ {
				fmt.Println("Index:", i, j, k, "--->", arr2[i][j][k])
			}
		}
	}
	arr3 := arr1
	arr3[0] = 1000 // This does not change arr1 because both of them are two different copies. Aka values types
	fmt.Println("Array-3", arr3, "Array-1", arr1)
	//arr4 := [4]int{0, 0, 0, 0}
	//var arr4 [4]int
	//var arr5 [3]int = arr4
	// changeVal([3]int(arr4), 0, 10) // cannot convert

}

func changeVal(arr [3]int, index, value int) error {
	if index >= len(arr) {
		return errors.New("index is out of bounds")
	}

	arr[index] = value

	return nil
}
