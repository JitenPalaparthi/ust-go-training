package rect

import "fmt"

type Rect struct {
	Length, Width   float32
	Area, Perimeter float32
}

// if the name of the type, its fields and methods start with Uppercase then they are
// exported types, fileds and methods.
// if not they are not exported so that they cannot be used other than in the same package

// methods and functions are different
// methods are created with receivers
// receiver is like normal parameter but it has to be given preceeded by method name.

// AreaOfRect is a method not a function because it has a receiver
func (r *Rect) AreaOfPtr() float32 {
	r.Area = r.Length * r.Width
	return r.Area
}

// PerimeterOfRect is a method not a function because it has a receiver
func (r *Rect) PerimeterOfPtr() float32 {
	r.Perimeter = 2 * (r.Length + r.Width)
	return r.Perimeter
}

func (r Rect) AreaOf() float32 {
	r.Area = r.Length * r.Width
	return r.Area
}

// PerimeterOfRect is a method not a function because it has a receiver
func (r Rect) PerimeterOf() float32 {
	r.Perimeter = 2 * (r.Length + r.Width)
	return r.Perimeter
}

func (r *Rect) ShowAreaAndPerimeter() {
	fmt.Printf("Area:%0.2f\n", r.Area)
	fmt.Printf("Perimeter:%0.2f\n", r.Perimeter)
}
