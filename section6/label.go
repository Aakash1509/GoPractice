package main

import "fmt"

func main() {
	people := [5]string{"Aakash", "Vraj", "Samarth", "Dhruv", "Keval"}
	friends := [2]string{"Vraj", "Marry"}

outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
			}
		}
		fmt.Println("Outer for loop")
	}
	fmt.Println("Next instruction after break")

	i := 0
loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}
}
