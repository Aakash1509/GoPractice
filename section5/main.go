package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		a int
		b float64
		c bool
		d string
	)

	x, y, z := 20, 15.5, "Gopher"

	_, _, _, _, _, _, _ = a, b, c, d, x, y, z

	//fmt.Println(a, b, c, d, x, y, z)

	const (
		seconds = 86400
		days    = 365
	)
	fmt.Printf("%v", seconds*days)

	const (
		jun = iota + 6
		july
		aug
	)
	fmt.Println(jun, july, aug)

	e, f, g := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	fmt.Printf("%d\n", e)
	fmt.Printf("%f\n", f)
	fmt.Printf("%s\n", g)
	fmt.Printf("%#v\n", score)
	fmt.Printf("%q\n", g)
	fmt.Printf("e is %v, f is %v, g is %v, score is %v\n", e, f, g, score)
	fmt.Printf("%T %T\n", f, score)

	var i int = 3
	i1 := float64(i)
	fmt.Printf("f1's type: %T, f1's value: %f\n", i1, i1)

	var h float64 = 3.2
	h1 := int(h)
	fmt.Println(h1)

	var i2 = 3
	var f2 = 3.2
	var s3, s4 = "3.14", "5"

	i3 := strconv.Itoa(i2)
	s4i, _ := strconv.Atoi(s4)
	sf2 := fmt.Sprintf("%f", f2)
	f3, _ := strconv.ParseFloat(s3, 64)
	fmt.Printf("%T,%v\n", i3, i3)
	fmt.Printf("%T,%v\n", s4i, s4i)
	fmt.Printf("%T,%v\n", sf2, sf2)
	fmt.Printf("%T,%v\n", f3, f3)

	type duration int

	var hour duration = 3600

	fmt.Println(hour)
	fmt.Printf("%T\n", hour)

	type mile float64
	type kilometer float64

	const m2km = 1.609

	var mileBerlinToParis mile = 655.3
	var kmBerlinToParis kilometer

	kmBerlinToParis = kilometer(mileBerlinToParis * m2km)
	fmt.Println(kmBerlinToParis)

}
