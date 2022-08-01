package shape

type Rect struct {
	Length, Width   float32
	Area, Perimeter float32
}

// AreaOfRect is a function
func AreaOfRect(rect Rect) float32 {
	rect.Area = rect.Length * rect.Width
	return rect.Area
}

// PerimeterOfRect is a function
func PerimeterOfRect(rect Rect) float32 {
	rect.Perimeter = 2 * (rect.Length + rect.Width)
	return rect.Perimeter
}

// AreaOfRect is a function
func AreaOfRectPtr(rect *Rect) float32 {
	rect.Area = rect.Length * rect.Width
	return rect.Area
}

// PerimeterOfRect is a function
func PerimeterOfRectPtr(rect *Rect) float32 {
	rect.Perimeter = 2 * (rect.Length + rect.Width)
	return rect.Perimeter
}
