package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var cities []string
	//cities = append(cities, "New York")
	//fmt.Println(cities)
	fmt.Println("cities is equal to nil: ", cities == nil)
	fmt.Printf("cities %#v\n", cities)
	fmt.Println(len(cities))

	//cities[0] = "London" will give runtime error
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	number := []int{2, 3, 4, 5, 6}
	fmt.Println(number)

	//s1 := "Go is cool!"
	//s1[0] = 'X' Error
	//fmt.Println(s1)

	nums := make([]int, 2)
	fmt.Printf("%#v\n", nums)
	fmt.Println(nums == nil)

	type names []string
	friends := names{"abc", "xyz"}
	fmt.Println(friends)

	//Iterating over slices
	for index, value := range numbers {
		fmt.Println("index ", index, "value ", value)
	}

	var n []int
	n = numbers
	fmt.Println(n)

	//Comparing slices

	var a []int
	fmt.Println(a == nil)

	b := []int{}
	fmt.Println(b == nil)

	//Ways to compare slices

	//1st iterating through for loop
	x, y := []int{1, 2, 3}, []int{1, 2, 3}

	//_, _ = a, b

	var eq bool = true

	x = nil //If I do this then eq will remain true , as code beneath this will not be executed
	for i, valueA := range x {
		if valueA != y[i] {
			eq = false
			break
		}
	}

	if len(x) != len(y) { //Therefore this is also necessary
		eq = false
	}

	if eq {
		fmt.Println("Slices are equal")
	} else {
		fmt.Println("Slices are not equal")
	}

	//Appending to slice
	numbers1 := []int{2, 3}

	numbers1 = append(numbers1, 20, 30, 40)
	fmt.Println(numbers1)

	//Appending a whole slice to another
	n1 := []int{100, 200}
	numbers1 = append(numbers1, n1...)
	fmt.Println(numbers1)

	//Copying a slice
	src := []int{10, 20, 30}
	//If size of dst is less it will copy till that , if greater it will copy 0 but count will be equal to elements copied
	//If dst has size 0 nothing will be copied
	dst := make([]int, len(src))
	nn := copy(dst, src) //number of elements copied
	fmt.Println(src, dst, nn)

	//Slice expressions
	aa := [5]int{1, 2, 3, 4, 5}
	bb := aa[1:3]
	fmt.Printf("%v,%T\n", bb, bb)

	//Appending and slicing
	s4 := []int{1, 2, 3, 4, 5, 6}
	s4 = append(s4[:4], 100)
	fmt.Println(s4)

	s4 = append(s4[:4], 200) //Will replace
	fmt.Println(s4)

	s2 := "中文维基是世界上"
	rs := []rune(s2)
	fmt.Println(string(rs[0:3]))

	tour := []string{"Aakash", "Sky", "Motadata"}
	fmt.Println(tour)

	for i, v := range tour {
		fmt.Println(i, v)
	}

	tour1 := []float64{1.1, 2.1, 3.1}

	tour1 = append(tour1, 10.1)

	fmt.Println(tour1)

	tour1 = append(tour1, 4.1, 5.5, 6, 6)
	fmt.Println(tour1)

	tour2 := []float64{5.1, 6.1}

	tour1 = append(tour1, tour2...)

	fmt.Println(tour1)

	args := os.Args[1:]

	sum := 0.
	product := 1.
	for _, v := range args {
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}
		sum += num
		product *= num
	}
	fmt.Printf("Sum : %v,Product : %v\n", sum, product)

	question := []int{5, -1, 9, 10, 1100, 6, -1, 6}

	question = question[2 : len(question)-2]

	fmt.Println(question)

	sum1 := 0

	for _, value := range question {
		fmt.Println(value)
		sum1 += value
	}
	fmt.Println(sum1)
}
