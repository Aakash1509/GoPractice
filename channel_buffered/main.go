package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	//spawning a go routine
	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			c <- i //After sending 4 it will be blocked as channel is full
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}
		close(c) //after sending all data closing channel
	}(c)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(5 * time.Second)

	for v := range c { //v:=<-c
		fmt.Println("main goroutine received value from channel: ", v)
	}

	//c<-10 : panic as channel is closed
	//<-c : ZERO value

}
