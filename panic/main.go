package main

import "fmt"

func handlePanic() {
	a := recover()
	if a != nil {
		fmt.Println("RECOVERED", a)
	}
}
func main() {
	fmt.Println(1)
	//panic("Error")
	runPanic()
	fmt.Println(2)

}

func runPanic() {
	defer handlePanic()
	panic("Critical error")
}
