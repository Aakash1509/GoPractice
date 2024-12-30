package main

import "fmt"

func main() {
	//1st

	type duration int

	var hour duration

	fmt.Printf("Type of hour %T and value %v\n", hour, hour)

	hour = 3600

	fmt.Printf("Type of hour %T and value %v\n", hour, hour)

	//3rd
	type mile float64
	type kilometer float64

	const m2km = 1.609

	var mileBerlinToParis mile = 655.3
	var kmBerlinToParis kilometer

	kmBerlinToParis = kilometer(mileBerlinToParis * m2km)

	fmt.Printf("%f", kmBerlinToParis)
}
