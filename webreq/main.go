package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Learning web service")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error getting in GET response", err)
		return
	}
	defer res.Body.Close()

	fmt.Printf("Type of response: %T\n", res)

	//Reading response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body", err)
		return
	}
	fmt.Println("response: ", string(data))
}
