package interfaces

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func Geo(s Shape) {
	fmt.Println(s.Area())
}

func GeoTest() {
	r := Rectangle{width: 42, height: 55}
	Geo(r)
	c := Circle{radius: 2}
	Geo(c)
}
