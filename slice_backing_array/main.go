package main

import (
	"fmt"
	"unsafe"
)

func main() {

	arr1 := [4]int{1, 2, 3, 4}

	arr2 := arr1

	arr2[0] = 10

	fmt.Println(arr1) //Unchangee

	slice1, slice2 := arr1[0:2], arr1[1:3]

	//arr1[1] = 50
	slice1[1] = 50

	fmt.Println(arr1, slice1, slice2) //Changed

	s1 := []int{1, 2, 3, 4, 5}

	s2 := s1

	s2[0] = 5

	fmt.Println(s1) //Changed

	//To create a complete new slice which has different backing array use "append function". So modifying 1st list won't affect another
	cars := []string{"Ford", "BMW", "Audi"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)
	cars[0] = "Nissan"
	fmt.Println(cars, newCars)

	s5 := []int{1, 2, 3, 4, 5}
	newSlice := s5[0:3]
	fmt.Println(len(newSlice), cap(newSlice))

	newSlice = s5[2:5]
	fmt.Println(len(newSlice), cap(newSlice))

	fmt.Printf("%p %p \n", &s5, &newSlice)

	a := [5]int{1, 2, 3, 4, 5}
	s6 := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("array's size in bytes: %d\n", unsafe.Sizeof(a))
	fmt.Printf("slice size in bytes: %d\n", unsafe.Sizeof(s6))

}
