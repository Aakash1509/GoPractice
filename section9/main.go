package main

import "fmt"

func main() {
	var cities [2]string
	fmt.Printf("%#v\n", cities)

	grades := [3]float64{5.5, 6, 6}
	fmt.Printf("%#v\n", grades)

	taskDone := [...]bool{true, false, true, true}
	fmt.Printf("%#v\n", taskDone)

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	for i, j := range grades {
		fmt.Println(i, j)
	}

	nums := [...]int{30, -1, -6, 90, -6}

	for _, v := range nums {
		if v > 0 && v%2 == 0 {
			fmt.Println(v, "is even")
		}
	}

}
