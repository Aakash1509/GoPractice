package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//for loop as while loop
	j := 10
	for j >= 0 {
		fmt.Println(j)
		j--
	}

	// handling of multiple variables in a for loop
	for i, j := 0, 100; i < 10; i, j = i+1, j+1 {
		fmt.Printf("i = %v, j = %v\n", i, j)
	}

	count := 0
	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13\n", i)
			count++
		}
		if count == 10 {
			break
		}
	}
	fmt.Println("Message after for loop")
}
