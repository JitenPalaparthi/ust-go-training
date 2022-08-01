package rect

import (
	"errors"
	"fmt"
)

type Rect struct {
	Length, Width float32
}

func New(l, w float32) (*Rect, error) {
	if l == 0 || w == 0 {
		return nil, errors.New("invalid input")
	}
	return &Rect{Length: l, Width: w}, nil
}

func (r *Rect) Area() float32 {
	return r.Length * r.Width
}

func (r *Rect) ShowArea() {
	fmt.Println("Area", r.Area())
}
