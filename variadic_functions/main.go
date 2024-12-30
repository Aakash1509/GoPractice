package main

import (
	"fmt"
	"strings"
)

func f1(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}
func f2(a ...int) {
	a[0] = 100
}

func SumAndProduct(a ...float64) (float64, float64) {
	sum := 0.
	prod := 1.
	for _, v := range a {
		sum += v
		prod *= v
	}
	return sum, prod
}

// Mixture of non-variadic and variadic arguments

// If the function takes parameters of different types, only the last parameter of a function can be variadic.
func personInformation(age int, names ...string) string {
	fullName := strings.Join(names, " ")
	returnString := fmt.Sprintf("Age:%d,name:%s", age, fullName)
	return returnString
}
func main() {
	f1(1, 2)
	f1()

	nums := []int{1, 2, 3}
	nums = append(nums, 4, 5, 6)
	f1(nums...) //Passing slice to variadic function

	f2(nums...)
	fmt.Println("Nums:", nums)

	s, p := SumAndProduct(2.0, 5., 10.)
	fmt.Println("Sum and product:", s, p)

	info := personInformation(21, "Aakash", "Hiren", "Saraiya")
	fmt.Println(info)

}
