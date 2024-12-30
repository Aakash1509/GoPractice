package main

import "fmt"

func foo() {
	fmt.Println("in foo")
}
func bar() {
	fmt.Println("in bar")
}

func foobar() {
	fmt.Println("in foobar")
}

func main() {
	defer foo()
	fmt.Println("After deferring foo")
	bar()
	defer foobar()

}
