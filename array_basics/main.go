package main

import "fmt"

func main() {
	defer func() { fmt.Println("function1") }()
	defer func() { fmt.Println("function2") }()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("function6")
		}
	}()
	defer func() {
		fmt.Println("function3")
		panic("function4")
	}()
	defer func() { fmt.Println("function5") }()

}
