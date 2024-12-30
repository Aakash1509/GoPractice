package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	welcome := "Welcome to homepage"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating: ")

	rating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", rating)

	final, err := strconv.ParseFloat(rating, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(final)
}
