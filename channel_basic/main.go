package main

import "fmt"

func main() {
	var ch chan int
	fmt.Println(ch)

	ch = make(chan int) //Stores address like pointer
	fmt.Println(ch)

	c := make(chan int) //declaration and initialization (2-way channel)

	//<- channel operator

	//SEND operation

	c <- 10

	//RECEIVE

	num := <-c

	fmt.Println(<-c)

	_ = num

	close(c) //Closing a channel
}
