package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition...")

	var wg sync.WaitGroup

	var mut sync.Mutex

	wg.Add(3)

	var score = []int{0}

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()
		fmt.Println("One W")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
	}(&wg, &mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Two W")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
	}(&wg, &mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Three W")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
	}(&wg, &mut)

	wg.Wait()

	fmt.Println(score)
}
