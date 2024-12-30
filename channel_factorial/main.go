package main

import (
	"fmt"
	"runtime"
	"time"
)

func factorial(n int, c chan int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	time.Sleep(time.Second * 10)
	c <- f
}

func main() {
	ch := make(chan int)

	//defer close(ch)

	go factorial(5, ch)
	fmt.Println("No. of goroutines :", runtime.NumGoroutine()) //2

	//fmt.Println("Hello before receving")
	f := <-ch
	//fmt.Println("Hello after receving")
	fmt.Println(f)
	//fmt.Println("Hello after receving")
	//for i := 1; i <= 20; i++ {
	//	fmt.Println("No. of goroutines after entering for loop :", runtime.NumGoroutine()) //1
	//	go factorial(i, ch)
	//	fmt.Println("No. of goroutines after go factorial:", runtime.NumGoroutine()) //2
	//	if i == 20 {
	//		break
	//	}
	//	f = <-ch
	//	fmt.Println(i, " ", f)
	//}
	//fmt.Println("After for loop")
	//f = <-ch
	//fmt.Println(f)

	//The channel automatically synchronizes access to the data being sent and received. For example, if multiple goroutines are sending data to the same channel, Go ensures that only one sender can send data at a time. Similarly, if multiple goroutines are receiving from the same channel, Go ensures that only one receiver can receive at a time.

	//Using anonymous function

	fmt.Println("---------------------------------------")
	for i := 5; i <= 15; i++ {
		fmt.Println("No. of goroutines after entering for loop :", runtime.NumGoroutine()) //1
		go func(n int, c chan int) {
			f := 1
			for i := 2; i <= n; i++ {
				f *= i
			}
			//time.Sleep(time.Second * 10)
			c <- f
		}(i, ch)
		fmt.Println("No. of goroutines after go anonymous function:", runtime.NumGoroutine()) //2
		fmt.Printf("Factorial of %d is %d\n", i, <-ch)
	}
}
