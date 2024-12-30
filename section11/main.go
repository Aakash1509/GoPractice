package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	for i, v := range s {
		fmt.Println(i, v)
	}
	s = append(s, 4, 5)
	fmt.Println(s)

	friends := []string{"Marry", "John", "Paul", "Diana"}
	friends1 := make([]string, len(friends))
	copy(friends1, friends)
	fmt.Println(friends1)

	str := "你好 Go!"

	runes := []rune(str)
	fmt.Println(string(runes))

	for i, r := range runes {
		fmt.Println(i, string(r))
	}

}
