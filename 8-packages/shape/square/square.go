package square

// There are not access modifiers /specifiers in go
// anything starts with UpperCase can be exported (similar to public)
// anything starts with lowerCase cannot be exported (similar to private)
import "fmt"

var (
	Side float32
)

// Area can be exported
func Area() float32 {
	return Side * Side
}

// Parimter can be exported
func Perimeter() float32 {
	return 4 * Side
}

// display cannot be exported
func display() {
	fmt.Println("This is square package")
}
