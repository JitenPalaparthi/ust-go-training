package main

import (
	"fmt"
	"shapes/shape"
	"strconv"
)

func main() {

	rect := shape.Rect{Length: 12.456, Width: 15.567}
	fmt.Println("Area and Perimeter through call by value functions")
	area1 := shape.AreaOfRect(rect)
	perimeter1 := shape.PerimeterOfRect(rect)
	fmt.Printf("Area of Rect:%.2f\n", area1)
	fmt.Printf("Perimter of Rect:%.2f\n", perimeter1)
	fmt.Println("Area of Rect from Rect Field:", rect.Area)
	fmt.Println("Perimeter of Rect from Rect Field:", rect.Perimeter)

	// Passing address of rect to a different area and perimeter functions that accesspts
	// pointer of rect
	fmt.Println("\n\nArea and Perimeter through call by reference functions")
	area2 := shape.AreaOfRectPtr(&rect)
	// Using formatter from strconv package
	areastr2 := strconv.FormatFloat(float64(area2), 'f', 2, 32)
	fmt.Println("Format Area of Rect using formatter", areastr2)
	perimeter2 := shape.PerimeterOfRectPtr(&rect)
	fmt.Printf("Area of Rect:%.2f\n", area2)
	fmt.Printf("Perimter of Rect:%.2f\n", perimeter2)
	fmt.Println("Area of Rect from Rect Field:", rect.Area)
	fmt.Println("Perimeter of Rect from Rect Field:", rect.Perimeter)
}
