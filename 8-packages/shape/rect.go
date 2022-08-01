package shape

// There are not access modifiers /specifiers in go
// anything starts with UpperCase can be exported (similar to public)
// anything starts with lowerCase cannot be exported (similar to private)
import "fmt"

var (
	Length, Width float32
)

// First time when the package is called

func init() {
	Length = 1
	Width = 1
}
func init() {
	display()
}

////  Area can be exported
// func Area() float32 {
// 	return Length * Width
// }

// // Parimter can be exported
// func Perimeter() float32 {
// 	return 2 * (Length + Width)
// }

// display cannot be exported
func display() {
	fmt.Println("This is shape package.Values of Length and Width", Length, Width)
}
