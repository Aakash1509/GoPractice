package main

import "fmt"

func main() {
	var c1 chan float64
	_ = c1

	c2 := make(<-chan rune)
	_ = c2

	c3 := make(chan<- rune)
	_ = c3

	c4 := make(chan int, 10)
	_ = c4

	fmt.Printf("%T\n", c1)
	fmt.Printf("%T\n", c2)
	fmt.Printf("%T\n", c3)
	fmt.Printf("%T\n", c4)

	ch := make(chan string)

	go func(n string) {
		ch <- n
	}("Gophers!")

	value := <-ch
	fmt.Println("Value received from channel:", value)
}
