package main

import (
	"fmt"
	"src/mypackages/numbers"
)

func main() {
	var x uint = 40
	fmt.Printf("%d is even: %t\n", x, numbers.Even(x))
}
