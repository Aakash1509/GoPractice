package main

import "fmt"

func main() {
	const x float64 = 5 //typed constant
	const y = 6         //untyped constant

	/*
		const p int = 4
		const y float64 = 2.2 * p
		This will produce an error as p is declared as int and is typed constant
	*/

	const p = 4 //It's an int but is not Go type of int , freedom is with us thus can multiply int with float
	const z = 2.2 * p
	fmt.Println("%T\n", z)

	//No need to typecast will automatically convert
	var i int = p
	var j float64 = p //Internal working : var j float64 = float64(p)
	var k byte = p
	fmt.Println(i, j, k)

	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	const (
		a1 = iota
		a2
		a3
	)

	fmt.Println(c1, c2, c3)

	const (
		a = iota * 2 //0
		b            //2
		c            //4
	)
	fmt.Println(a, b, c)

	//e=-2,f=-4,g=-5  : Skipping value  by blank identifier
	const (
		e = -(iota + 2)
		_
		f
		g
	)
	fmt.Println(e, f, g)
}
