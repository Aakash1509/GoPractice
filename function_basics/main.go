package main

import (
	"fmt"
	"math"
)

func f1() {
	fmt.Println("hello world")
}

func f2(a int, b int) {
	fmt.Println("Sum: ", a+b)
}

// shorthand parameters function
func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}
func f4(a float64) float64 {
	return math.Pow(a, a)
	//any statements below the return statement are never executed

}

func f5(a, b int) (int, int) {
	return a + b, a * b
}

// defining a function that have one parameter of type int and returns a "named parameter"
func sum(a, b int) (s int) {
	s = a + b //s is a variable with the zero value inside the function
	return    //Naked return functions (use only in short functions). Automatically returns
}

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	f1()
	f2(5, 7)
	f3(4, 5, 6, 4.4, 7.8, "golang")
	p := f4(4.1)
	fmt.Println(p)
	a, b := f5(8, 9)
	fmt.Println(a, b)
	mySum := sum(a, b)
	fmt.Println(mySum)

	//Anonymous functions
	func(msg string) {
		fmt.Println(msg)
	}("I am anonymous string")

	z := increment(10)
	fmt.Printf("%T\n", z)
	z()
	fmt.Println(z())

	func(vara int) {
		fmt.Println(vara)
	}(6)
}
