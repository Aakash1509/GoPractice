package main

import "fmt"

func main() {
	name := "Aakash"
	fmt.Println(&name)

	var x int = 2
	ptr := &x

	fmt.Printf("ptr is of type %T with value of %v and address of pointer %p\n", ptr, ptr, &ptr)
	fmt.Printf("address of x is %p\n", &x)

	//Declaring pointer of float64 type
	var ptr1 *float64 //zero value is nil
	_ = ptr1

	// creating a pointer using new() built-in function.
	p := new(int) // it creates a pointer called p that is a pointer to an int type
	x = 100
	p = &x

	fmt.Printf("p is of type %T with value of %v\n", p, p)

	*p = 90 //equivalent to x=90
	fmt.Println(x, *p)

	// In a nutshell:
	// &value => pointer -> if you have a value you turn it into an address or pointer by using the ampersand operator
	// *pointer => value  -> and if you have pointer you turn it into value by using the star operator

	a := 5.5
	p1 := &a
	pp1 := &p1
	fmt.Printf("Value of p1:%v,address of p1:%v\n", p1, &p1)
	//fmt.Printf("Value of p1:%v,address of p1:%p\n", p1, &p1) //same
	fmt.Printf("Value of pp1:%v,address of pp1:%v\n", pp1, &pp1)

	fmt.Printf("*p1 is %v\n", *p1)
	fmt.Printf("*pp1 is %v\n", *pp1)
	fmt.Printf("**pp1 is %v\n", **pp1)

	**pp1++
	fmt.Printf("a is %v\n", a)

	//Comparing pointers

	var p2 *int
	fmt.Printf("%#v\n", p2)

	y := 5
	p2 = &y

	z := 5
	p3 := &z

	fmt.Println(p2 == p3)

	p4 := &y
	fmt.Println(p2 == p4)

	c := 5
	p10 := &c
	*p10 = *p10 * 2
	//fmt.Println(*p10 + 1)
	fmt.Println(c)
}
