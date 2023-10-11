package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{
		Width:  2,
		Height: 3,
	}

	fmt.Println("Original Circle:")
	PrintShapeProperties(c)

	fmt.Println("Original Rectangle:")
	PrintShapeProperties(r)

}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Func2() {

}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func PrintShapeProperties(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())
}
