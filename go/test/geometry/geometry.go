package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	if r.Width < 0 || r.Height < 0 {
		return 0
	}
	return r.Height * r.Width
}

func main() {
	r := Rectangle{Width: 10, Height: 20}
	area := r.Area()
	fmt.Println(area)
}
