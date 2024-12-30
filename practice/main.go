package main

import (
	"fmt"
	"time"
)

func f1() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}

func f2(n int, ch chan int) {
	time.Sleep(5 * time.Second)
	ch <- n
}
func main() {
	fmt.Println("main execution started")
	go f1()
	c := make(chan int)

	go f2(10, c)

	n := <-c
	fmt.Println(n)

	fmt.Println("main execution ended")
}
