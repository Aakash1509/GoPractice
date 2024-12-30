package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}
func printValue(s shape) {
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}
func main() {
	c1 := circle{radius: 5}
	printValue(c1)
	r1 := rect{width: 3, height: 4}
	printValue(r1)

	var s shape = circle{radius: 5}

	fmt.Printf("Sphere Volume:%v\n", s.(circle).volume())

	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type\n", value)
	case rect:
		fmt.Printf("%#v has rectangle type\n", value)
	}

	//a := uint(8)
	//a=int(a)
	slice := []interface{}{1, "abc", "def"}
	fmt.Println(slice)

	change(slice)
	fmt.Println(slice)
}

func change(slice []interface{}) {
	slice = append(slice, "ehj")
}
