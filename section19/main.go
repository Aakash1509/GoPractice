package main

import "fmt"

func main() {
	type person struct {
		name   string
		age    int
		colors []string
	}

	me := person{name: "Aakash", age: 21, colors: []string{"red", "blue"}}

	you := person{name: "Sky", age: 22, colors: []string{"red", "blue"}}
	fmt.Printf("me: %+v\n", me)
	fmt.Printf("you: %+v\n", you)

	me.name = "Andrei"

	val := you.colors
	fmt.Printf("The type is %T and val is %v", val, val)

	val = append(val, "green")
	you.colors = val
	fmt.Printf("The type is %T and val is %v", you.colors, you.colors)

	for _, v := range me.colors {
		fmt.Println(v)
	}
}
