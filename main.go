package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

// go has a capability to return multiple values from a function
func swap(x, y string) (string, string) {
	return y, x
}

// implicit return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// explicit return
func splitt(sum int) (int, int) {
	x := sum * 4 / 9
	y := sum - x
	return x, y
}

func main() {
	fmt.Println("Hello", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(add(7, 11))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	fmt.Println(splitt(17))
}
