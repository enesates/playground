package main

import (
	"fmt"
	"math"
)

type Measurable interface {
	Area() float64
	Perimeter() float64
}

/////////

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

/////////

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

/////////

type Triangle struct {
	A, B, C float64
}

func (t Triangle) Area() float64 {
	return t.A * t.B * 0.5
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

/////////

func TotalArea(shapes []Measurable) float64 {
	totalArea := 0.0

	for _, shape := range shapes {
		totalArea += shape.Area()
	}

	return totalArea
}

func TotalPerimeter(shapes []Measurable) float64 {
	totalPerimeter := 0.0

	for _, shape := range shapes {
		totalPerimeter += shape.Perimeter()
	}

	return totalPerimeter
}

func main() {
	shapes := []Measurable{
		Rectangle{Width: 2.1, Height: 4.23},
		Circle{Radius: 4.5},
		Triangle{A: 3, B: 4, C: 5},
	}

	fmt.Printf("Total area: %.2f\n", TotalArea(shapes))
	fmt.Printf("Total perimeter: %.2f\n", TotalPerimeter(shapes))
}
