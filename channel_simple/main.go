package main

import "fmt"

func f1(n int, c chan int) {
	c <- n
}
func main() {
	ch := make(chan int)

	//only for receiving
	c1 := make(<-chan string)

	//only for sending
	c2 := make(chan<- string)

	fmt.Printf("%T, %T, %T\n", ch, c1, c2)

	go f1(10, ch) //Creating go routine

	n := <-ch

	fmt.Println("Received value of n : ", n)
	fmt.Println("Main exit")

}
