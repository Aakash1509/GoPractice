package main

import "fmt"

func main() {
	//1st way
	//var i int = 5
	var i = 5
	fmt.Println("Value of i is: ", i)

	//If variable is not used
	var name = 5
	_ = name

	//2nd way
	money := 35.5
	money = 36.5
	fmt.Println("Value of money is: ", money)

	//Declaring Multiple variables

	//1st way
	car, cost := "Audi", 5000
	fmt.Println(car, cost)

	//While redeclaring atleast 1 variable should be new , else it will produce compile time error
	car, year := "BMW", 2024
	_ = year

	//2nd way
	var (
		salary    float64
		firstName string
		gender    bool
	)
	fmt.Println(salary, firstName, gender)

	//3rd way - Same datatype
	var a, b, c int
	fmt.Println(a, b, c)

	//Multiple assignments

	var p, j int
	p, j = 5, 8

	j, p = p, j //Swapping variables
	fmt.Println(p, j)

	sum := 5 + 2.3
	fmt.Println(sum)

	var r rune = 'f'
	fmt.Printf("%T\n", r)
	fmt.Printf("%c\n", r)

	//Compposite data types

	//Array type
	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers)

	//Slice type
	var cities = []string{"London", "Tokyo", "New York"}
	fmt.Printf("%T\n", cities)

	//Map type
	balances := map[string]float64{
		"USD": 565,
		"EUR": 511.2,
	}
	fmt.Printf("%T\n", balances)

	//Struct type
	type Person struct {
		name string
		age  int
	}
	var you Person
	fmt.Printf("%T\n", you)

	//Pointer type
	var x int = 2
	ptr := &x //Stores address of x
	fmt.Printf("ptr is of type %T with value of %v\n", ptr, ptr)

	//Function type
	fmt.Printf("%T\n", f)
}
func f() {

}
