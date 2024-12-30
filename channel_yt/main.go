package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Figuring out channel")

	mychan := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)

	go func(mychan chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println(<-mychan)
	}(mychan, &wg)
	go func(mychan chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		mychan <- 5
	}(mychan, &wg)

	wg.Wait()
}
