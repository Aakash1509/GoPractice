package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	case t.Hour() < 19:
		fmt.Println("Good evening!")
	default:
		fmt.Println("Good night!")
	}
	fmt.Print("Go runs on ")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}
