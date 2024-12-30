package main

import (
	"fmt"
	"math"
)

// an interface contains only the signatures of the methods, but not their implementation
type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

//func printCircle(c circle) {
//	fmt.Println("Shape: ", c)
//	fmt.Println("Area: ", c.area())
//	fmt.Println("Perimeter: ", c.perimeter())
//}
//
//func printRectangle(r rectangle) {
//	fmt.Println("Shape: ", r)
//	fmt.Println("Area: ", r.area())
//	fmt.Println("Perimeter: ", r.perimeter())
//}

// any type that implements the interface is also of type of the interface
// rectangle and circle values are also of type shape
func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())

}
func main() {

	c1 := circle{radius: 5.}
	r1 := rectangle{width: 10., height: 5.}
	//printCircle(c1)
	//printRectangle(r1)
	print(c1)
	print(r1)

	var s shape
	fmt.Printf("%T\n", s)

	ball := circle{radius: 5}
	s = ball

	print(s)
	fmt.Printf("Type of s : %T\n", s)

	var s1 shape = circle{2.5}
	fmt.Printf("%T\n", s1)

	//s1.volume() //this of type circle but still will give error as interface doesn't have volume method

	//So to avoid this error we use type assertion
	//fmt.Println(s1.(circle).volume())

	//To check whether assertion succeeded or not
	ball1, ok := s1.(circle)
	if ok {
		fmt.Printf("Ball volume : %v\n", ball1.volume())
	}

	//** TYPE SWITCHES **//

	// it permits several type assertions in series
	s1 = rectangle{width: 10., height: 5.}
	switch value := s1.(type) {
	case circle:
		fmt.Printf("%#v has circle type", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type", value)
	}
}
