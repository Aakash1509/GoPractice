package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v", numbers)

	n1 := []int{}
	n1 = append(n1, numbers...)
	fmt.Printf("%v\n", n1)

	var1 := 'a' //rune
	fmt.Printf("%T\n", var1)
	fmt.Printf("%v\n", var1)
	fmt.Printf("%c\n", var1)

	s1 := "Go is cool!"
	fmt.Println(s1[0:2])
}
