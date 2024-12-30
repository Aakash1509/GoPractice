package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//go greeter("Hello")
	//greeter("World")

	websites := []string{
		"https://golang.org",
		"https://google.com",
		"https://facebook.com",
		"https://github.com",
	}

	for _, website := range websites {
		go getStatusCode(website, &wg)
		wg.Add(1)
	}

	wg.Wait()

}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		fmt.Printf("%d status code for %s\n", resp.StatusCode, url)
	}
}
