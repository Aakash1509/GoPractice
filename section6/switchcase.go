package main

import (
	"fmt"
	"time"
)

func main() {
	language := "golang"

	switch language {
	case "go", "golang":
		fmt.Println("You are learning Go language")
	case "Python":
		fmt.Println("You are learning Python language")
	default:
		fmt.Println("You are not learning other language")
	}

	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
