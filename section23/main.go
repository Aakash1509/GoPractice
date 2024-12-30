package main

import "fmt"

func swap(a, b *float64) {
	*a, *b = *b, *a
}

func main() {
	x, y := 10, 2
	ptrx, ptry := &x, &y

	z := *ptrx / *ptry
	fmt.Println(z)

	a, b := 5.5, 8.8
	swap(&a, &b)
	fmt.Println(a, b)

}
