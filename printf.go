package main

import "fmt"

func main() {
	fmt.Printf("Hello world %d\n", 21)
	fmt.Printf("x is %d,y is %f\n", 5, 6.8)
	fmt.Printf("He says: \"Hello Go!\"\n")

	x := 3.4
	y := 6.9
	fmt.Printf("%#v\n", x)
	fmt.Printf("x*y=%.3f\n", x*y)
}
