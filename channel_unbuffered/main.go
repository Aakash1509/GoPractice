package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int) //unbuffered channel

	c2 := make(chan int, 3) //buffered channel
	_ = c2

	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into channel")
		c <- 10 //blocking till main doesn't read data from channel

		//The following statement will be only printed after main goroutine will receive the value
		fmt.Println("Hello")
		fmt.Println("func goroutine after sending data into channel")
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(2 * time.Second)

	fmt.Println("main goroutine starts receiving data from channel")
	d := <-c1
	fmt.Println("main goroutine received data:", d)

	time.Sleep(time.Second)

}
