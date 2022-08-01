package main

import "fmt"

func main() {
	greet()
	greetBy("Hello World!")
	greetByName("Jiten!", "Hello")
	msg := greetReturn()
	fmt.Println(msg)

	a1, p1 := rectAreaAndPerimeter(12.45, 15.67)

	fmt.Println("Area a1", a1, "Perimeter p1", p1)

	a2, _ := rectAreaAndPerimeter(12.45, 15.67) // _ blank identifier

	fmt.Println("Area a2", a2)

	_, p2 := rectAreaAndPerimeter(12.45, 15.67)

	fmt.Println("Perimeter p2", p2)

	l := 12.45                                   // float64
	w := 15.67                                   // float64
	rectAreaAndPerimeter(float32(l), float32(w)) // by default float64 cannot be used as a arg for float32

	num := 100

	fmt.Println("Num before calling changeVal", num)

	changeVal(num)

	fmt.Println("Num after calling changeVal", num)

}

// func identifier(params)(returns){}
func greet() {
	fmt.Println("Hello World!")
}

func greetBy(message string) {
	fmt.Println(message)
}

func greetByName(message, name string) {
	fmt.Println(message, name)
}

// func string greetReturn() //Not OK

func greetReturn() string { // string is not a named parameter
	return "Hello World!"
}

func rectAreaAndPerimeter(l, w float32) (area float32, perimeter float32) { // this func returns two parameters.
	//a and p are named parameters
	area = l * w
	perimeter = 2 * (l + w)
	//return l * w, 2 * (l + w) OK
	return area, perimeter
}

func changeVal(i int) {
	i = 200
}
