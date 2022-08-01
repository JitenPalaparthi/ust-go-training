package main

// There are 3 kinds of packages
// 1- standard packages: eg. fmt,math,net,time,strconv,strings,bufio,io
// 2- user defined packages : Packages that reside inside the code, usually created by the same author
// 3- third party packages :  eg: github.com/labstack/echo/v4 . downloaded from github or similar systems. Usually these are created by others

import (
	"fmt"
	"math/rand" // This is a package inside another package
	"demo/shape" // go mod name becomes root path for the user defined package
	"demo/shape/square"
)

func main() {

	greet()
	fmt.Println("Generate random number", rand.Intn(100))

	shape.Length = 12.45
	shape.Width = 15.67

	fmt.Println("Area of Rect", shape.Area())
	fmt.Println("Perimeter of Rect", shape.Perimeter())

	square.Side = 25.25

	fmt.Println("Area of Square", square.Area())
	fmt.Println("Perimeter of Square", square.Perimeter())

}
