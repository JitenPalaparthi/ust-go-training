package main

import (
	"fmt"
	"shapes/shape/rect"
)

func main() {

	fmt.Println("The rect varialbe is not a pointer by creation")
	rect1 := rect.Rect{Length: 12.45, Width: 15.567}
	fmt.Println("calling method which is a non pointer receivers")
	rect1.AreaOf()
	rect1.PerimeterOf()
	rect1.ShowAreaAndPerimeter()

	fmt.Println("calling method which is a pointer receivers")
	rect1.AreaOfPtr()
	rect1.PerimeterOfPtr()
	rect1.ShowAreaAndPerimeter()

	fmt.Println("---------------------------------")
	fmt.Println("The rect varialbe is a pointer by creation")

	rect2 := &rect.Rect{Length: 12.45, Width: 15.567}
	fmt.Println("calling method which is a non pointer receivers")
	rect2.AreaOf()
	rect2.PerimeterOf()
	rect2.ShowAreaAndPerimeter()

	fmt.Println("calling method which is a pointer receivers")
	rect2.AreaOfPtr()
	rect2.PerimeterOfPtr()
	rect2.ShowAreaAndPerimeter()

}
