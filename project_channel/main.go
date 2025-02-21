package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func checkAndSaveBody(url string, c chan string) {
	resp, err := http.Get(url)

	if err != nil {
		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error: %v\n", err)
		c <- s //sending info
	} else {
		defer resp.Body.Close()
		s := fmt.Sprintf("%s -> Status Code: %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := io.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1] //E.g : http://www.google.com it will store fileName as www.google.com
			file += ".txt"
			s += fmt.Sprintf("Writing response body to %s\n", file)

			err = os.WriteFile(file, bodyBytes, 0644)
			if err != nil {
				//log.Fatal(err)
				s += "Error writing file\n"
				c <- s
			}
			s += fmt.Sprintf("%s is UP\n", url)
			c <- s
		}
	}
}

func main() {

	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	//1: Make a channel
	c := make(chan string)

	for _, url := range urls {
		go checkAndSaveBody(url, c)
		fmt.Println(strings.Repeat("-", 20))
	}
	fmt.Println("No of goroutines: ", runtime.NumGoroutine())

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

}
