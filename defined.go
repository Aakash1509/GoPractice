package main

import "fmt"

type km float64
type mile float64

type minute uint
type duration minute

func main() {

	type age int        //int is its underlying type
	type oldAge age     //int is its underlying type
	type veryOldAge age //int is its underlying type

	type speed uint //uint is its underlying datatype

	var d duration
	fmt.Printf("%T\n", d)

	var s1 speed = 10
	var s2 speed = 20

	fmt.Println(s2 - s1)

	var x uint
	//x=s1 //error : different types
	x = uint(s1)
	_ = x

	var parisToLondon km = 465
	var distanceInMiles mile
	distanceInMiles = mile(parisToLondon) / 0.621
	fmt.Println(distanceInMiles)

	//Alias

	var a uint8 = 10
	var b byte

	b = a //No error as byte and uint8 are alias
	_ = b
	type second = uint //Created alias

	var hour second = 3600
	fmt.Printf("Minutes in an hour: %d\n", hour/60) // => hour type: uint
	//no need to convert operations (same type)

}
