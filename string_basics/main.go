package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Hello"
	fmt.Println(s1)

	//Using double-quotes
	fmt.Println("He say : \"Hello\"")

	//2nd way , using backtick : Raw string
	fmt.Println(`He say : "Hello !"`)

	//Multiline string
	fmt.Println("Price:100\nBrand:Bike")

	fmt.Println(`
Price:100
Brand:Bike`)

	//Using backslashes within a string
	fmt.Println(`C:\Users\Andrei`)
	fmt.Println("C:\\Users\\Andrei")

	//Strings are immutable
	var s3 = "I love " + "Go" + "Programming"
	fmt.Println(s3 + "!")

	//Getting an element of string
	fmt.Println("Element at index zero: ", s3[0]) //It will print ASCII as character is not supported

	str := "tara"
	fmt.Println(len(str))
	fmt.Println(str[1])

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	str1 := "ţară"
	fmt.Println(len(str1))

	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c\n", str1[i])
	}

	for i := 0; i < len(str1); {
		r, size := utf8.DecodeRuneInString(str1[i:])
		fmt.Printf("%c\n", r)
		i += size
	}

	for _, r := range str1 {
		fmt.Printf("%c\n", r)
	}

	m := utf8.RuneCountInString(str1)
	fmt.Println(m)

	//Slicing a rune
	rs := []rune(str1)
	fmt.Printf("%T\n", rs)
	fmt.Printf(string(rs[0:3]))

}
