package shape

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	length float64
	width  float64
}

func NewRectangle(length float64, width float64) Shape {
	return Rectangle{length: length, width: width}
}
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

type Circle struct {
	radius float64
}

func NewCircle(radius float64) Shape {
	return Circle{radius: radius}
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
