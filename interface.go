package main

import "fmt"

// import "fmt"

// Define inteface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// struct : Rectange
type Rectangle struct{
	Width, Height float64
}

// Rectnage implements Shape
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64{
	return 2 * (r.Width + r.Height)
}

// Circle

type Circle struct{
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius 
}

func PrintShapeInfo(s Shape) {
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter ", s.Perimeter())
}