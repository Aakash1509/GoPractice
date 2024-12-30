package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	sum := 1
	for i := 0; i < 10; i++ {
		sum += i
	}

	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum)
	fmt.Println(sum1)
	fmt.Println(sqrt(2), sqrt(-4))
}
