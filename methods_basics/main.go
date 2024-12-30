package main

import (
	"fmt"
)

type names []string

type aakash []string

func (n aakash) print1() {
	for i, name := range n {
		fmt.Println(i, name)
	}
}

// declaring a method (function receiver)
func (n names) print() {
	// n is called method's receiver
	// n is the actual copy of the names we're working with and is available in the function.
	// n is like this or self from OOP.
	// any variable of type names can call this function on itself like variable_name.print()
	for i, name := range n {
		fmt.Println(i, name)
	}
}

type car struct {
	brand string
	price int
}

func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice

	// the changes are not seen to the outside world (pass by value)
}

// declaring method with pointer receiver
func (c *car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand // equivalent to c.brand = newBrand
	(*c).price = newPrice

	// the changes are seen the outside world
}
func main() {
	friends := names{"abc", "efg"}

	friends1 := aakash{"hello", "Aakash"}
	friends.print() //equivalent to names.print(friends)
	friends1.print1()
	names.print(friends)

	//declaring car value
	myCar := car{brand: "Audi", price: 4000}
	myCar.changeCar1("Audi", 100)
	fmt.Println(myCar) //No change

	// calling the method with a pointer receiver. It changes the value!

	(&myCar).changeCar2("Ford", 3000) //Equivalent to myCar.changeCar2("Ford", 10000)
	fmt.Println(myCar)

	// declaring a pointer to car
	var yourCar *car
	yourCar = &myCar      // it stores the address of myCar
	fmt.Println(*yourCar) // -> {Seat 25000}

	//calling method on pointer type
	yourCar.changeCar2("VW", 30000)
	(*yourCar).changeCar2("VW", 30000)
	fmt.Println(*yourCar) // => {VW 30000}
}
