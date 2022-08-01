package main

import (
	"fmt"
	iface "shapes/interfaces"
	"shapes/shape/rect"
	"shapes/shape/square"
)

func main() {
	r1, err := rect.New(12.45, 15.67)
	s1, err := square.New(25.25)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		Area(r1)
		Area(s1)
	}

	shapes := make([]iface.IArea, 0) // what if you give interface{}
	shapes = append(shapes, &rect.Rect{Length: 12.56, Width: 15.12})
	shapes = append(shapes, &rect.Rect{Length: 34.76, Width: 55.67})
	shapes = append(shapes, &rect.Rect{Length: 16.89, Width: 18.98})
	shapes = append(shapes, &rect.Rect{Length: 13.45, Width: 9.5})
	shapes = append(shapes, &square.Square{Side: 12.25})
	shapes = append(shapes, &square.Square{Side: 44.75})
	shapes = append(shapes, &square.Square{Side: 32.64})
	shapes = append(shapes, &square.Square{Side: 21.89})

	for _, v := range shapes {
		Area(v)
	}

}

func Area(iarea iface.IArea) {
	fmt.Println("Area:", iarea.Area())
}
