package main

import "fmt"

func main() {
	m := map[int]bool{1: true, 2: false, 3: false}

	fmt.Println(m)
	delete(m, 1)
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	var m1 map[int]int
	fmt.Printf("Type of m1 is %T and value is %v\n", m1, m1)

	m2 := map[int]string{
		1: "Aakash",
		2: "Sky",
	}
	fmt.Println(m2)

	m2[10] = "Abba"
	fmt.Println(m2)
}
