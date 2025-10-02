package main

import (
	"fmt"
	
)

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

// Empty interface
func describe (i interface{}){
	fmt.Printf("Type = %T, Value = %v\n", i , i)
}



	// Type Assertion

	// var inter interface{} = "hello"

	// chk := inter.(string)
	// fmt.Println(chk)

	// chk1, ok := inter.(string)
	// fmt.Println(chk1, ok)

	// chk2, ok := inter.(float64)
	// fmt.Println(chk2, ok)


	// Type Switch
	
	func do(i interface {}){
		switch v := i.(type) {
		case int:
			fmt.Printf("Twice %v is %v\n" , v , v*2)
		case string:
			fmt.Printf("%q is %v bytes long\n", v , len(v))
		default:
			fmt.Printf("I dont know about type %T\n", v)
		}
	}


	// generics

	func Index[T comparable](s []T, v T) int {
		for i, x := range s {
			if x == v {return i}
		}
		return  -1
	}
