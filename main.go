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

// variables
var c, python, java bool

var (
	ToBe   bool   = false
	MaxInt uint64 = 1<<64 - 1
	str    string
)

// struct
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	z  = &Vertex{1, 2}
)

type Area struct {
	Lat, Long float64
}

// var m map[string]Area

func compute(fn func(float64, float64)float64) float64{
	return fn(3,4)
}

// Closure
func adder() func(int) int{
	sum := 0  // variable from outer fucntion
	return func(x int) int {
		sum += x // inner function uses outer variable
		return sum
	}
}

// Method
type Ex1 struct{
	X,Y float64
}

func (v Ex1) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


type person struct{
		Name string
		Age int
	}

	// method from Stringer interface
	func (p person) String() string  {
		return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
	}



func main() {
	fmt.Println("Hello", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(add(7, 11))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	fmt.Println(splitt(17))
	i := 10
	fmt.Println(c, python, java, i)
	fmt.Println(ToBe)
	fmt.Println(MaxInt)
	fmt.Printf("%q", str)

	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Pointer
	j, k := 42, 2701
	p := &j         // point to j
	fmt.Println(*p) // read  through pointer
	*p = 21
	fmt.Println(j)

	p = &k
	*p = *p / 37
	fmt.Println(k)

	fmt.Println(v1, v2, v3, z)

	// Arrays
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "World"
	fmt.Println(arr)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slice
	sl := primes[1:4]
	fmt.Println(sl)
	sl[0] = 1
	fmt.Println(sl)
	fmt.Println(primes)

	q := []int{2, 3, 4, 5, 6, 7, 8}
	fmt.Println(q)
	// var nums []int
	// for i := 0; i < 20; i++ {
	// 	nums = append(nums, i)
	// 	fmt.Printf("LengthL %d, CapacityL %d\n", len(nums), cap(nums))
	// }

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Println(i, v)
	}

	// for i := 0; i< len(pow); i++{
	// 	fmt.Printf("2**%d = %d\n", i, pow[i])
	// }
	// i -> index of the eleent in the slice
	// v -> value at that index

	m := make(map[string]Area)
	m["Bell Labs"] = Area{
		40.8, -74.3,
	}
	fmt.Println(m["Bell Labs"])

	map1 := make(map[string]int)
	map1["Answer"] = 47
	fmt.Println("The Value:", map1["Answer"])

	map1["Answer"] = 40
	fmt.Println("The Value:", map1["Answer"])

	delete(map1, "Answer")
	fmt.Println("The Value:", map1["Answer"])

	v, ok := map1["Answer"]
	fmt.Println("The Value", v, "Present?", ok)

	
	// Anonymous function
	hypot := func(x,y float64) float64{
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))

	posSum := adder()
	fmt.Println(posSum(1))
	fmt.Println(posSum(2))
	fmt.Println(posSum(3))

	v1 := Ex1{3,4}
	fmt.Println(v1.Abs())

	var s Shape
	s = Rectangle{Width: 4, Height: 5}
	fmt.Println("Rectangle Area:", s.Area())
	fmt.Println("Rectange Perimeter:", s.Perimeter())

	s = Circle{Radius: 3}
	fmt.Println("Circle Area", s.Area())
	fmt.Println("Circle Perimeter", s.Perimeter())

	r := Rectangle{7,8}
	c := Circle{7}

	PrintShapeInfo(r)
	PrintShapeInfo(c)

	describe(42)
	describe("Hello")
	describe(true)


	// Type Assertion
	var inter interface{} = "hello"

	chk := inter.(string)
	fmt.Println(chk)

	chk1, ok := inter.(string)
	fmt.Println(chk1, ok)

	chk2, ok := inter.(float64)
	fmt.Println(chk2, ok)

	do(2)
	do("hello")
	do(true)
	



	first := person{"sadat", 22}
	second := person{"sahil", 23}
	// fmt here checks if the type implements the fmt.Stringer interface
	fmt.Println(first, second)

	newI := Index([]int {10,20,30,40}, 20)
	fmt.Println(newI)
}
