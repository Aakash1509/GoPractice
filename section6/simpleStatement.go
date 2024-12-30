package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if args := os.Args; len(args) != 2 {
		fmt.Println("One argument is required")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("Argument must be an integer! Error", err)
	} else {
		fmt.Printf("%d km in miles is %v\n", km, float64(km)*0.621)
	}
}
