package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

const (
	Big   = 1 << 100
	small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

var c, python, java bool

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

var i, j = 1, 2

func add(a, b int) int {
	k := 5
	_ = k
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}

// Naked return
func example(a int) (ans int) {
	ans = a + 1
	return
}
func main() {
	//var i int
	var x, y = 3, 4
	z := 5.
	const Truth = true
	fmt.Println(needInt(small))
	fmt.Println(needFloat(small))
	fmt.Println(needFloat(Big))
	fmt.Println("Go rules?", Truth)
	fmt.Println()
	var f = math.Sqrt(float64(x*x + y*y))
	fmt.Println(x, y, f)
	fmt.Println(float64(x))
	fmt.Println(uint(z))
	fmt.Println(i, c, python, java)
	fmt.Println("My favourite number is ", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(add(1, 2))
	fmt.Println(swap("hello", "world"))
	fmt.Println(example(1000))
	fmt.Printf("%T %v\n", ToBe, ToBe)
	fmt.Printf("%T %v\n", MaxInt, MaxInt)
	fmt.Printf("%T %v\n", z, z)
}
