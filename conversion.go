package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Converting int to string
	s := string(99)
	fmt.Println(s)

	//Converting float to string
	s1 := fmt.Sprintf("%f", 44.2)
	fmt.Printf("Type is %T and string is %v\n", s1, s1)

	//Converting string to float
	s2 := "3.124"
	var f1, err = strconv.ParseFloat(s2, 64)
	_ = err
	fmt.Println(f1)

	i, err := strconv.Atoi("50")
	s3 := strconv.Itoa(20)
	fmt.Printf("i type is %T,i value is %v\n", i, i)
	fmt.Printf("s3 type is %T,s3 value is %q\n", s3, s3)
}
