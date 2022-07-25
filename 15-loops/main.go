//goto
package main

import "fmt"

func main() {

	var i int16 = 1
iterate:
	//{
	fmt.Print(i, " ")
	i++
	if i <= 10 {
		goto iterate
	}
	//}
	fmt.Println("\nGoto ends")

	fmt.Println("1-10 even numbers")

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Print(i, " ")
	}

	fmt.Println("\n1-10 odd numbers")

	for i := 1; ; i++ {
		if i > 10 {
			break
		}
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println("\n1-10 numbers")

	for i := 1; ; {
		if i > 10 {
			break
		}
		fmt.Print(i, " ")
		i++
	}
	fmt.Println("\nMore inits, conditions and increments")

	for i, j := 1, 2; i%j != 10; i, j = i+1, j+1 {
		fmt.Println("i--j", i, j)
	}

	/// print 1-100 prime numbers using above logic

	fmt.Println("Nested loop")

	n := 10
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// i,j -->if i==j break the nested loop
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("\ni=%d,j=%d", i, j)
			if i == j {
				break outer
			}
		}
	}
}
