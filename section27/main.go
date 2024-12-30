package main

import (
	"fmt"
	"math"
	"sync"
)

func sayHello(n string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Hello, %s!\n", n)
}

func sum(a, b float64, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%.2f\n", a+b)
}

func main() {
	var wg sync.WaitGroup

	//Question-1
	//wg.Add(1)
	//go sayHello("Mr. Wick", &wg)

	//Question-2
	//wg.Add(3)
	//go sum(10.1, 11.1, &wg)
	//go sum(-10.1, -10.1, &wg)
	//go sum(5.5, 4.4, &wg)

	//Question-3
	wg.Add(50)
	for i := 100; i < 150; i++ {
		go func(f int, wg *sync.WaitGroup) {
			defer wg.Done()
			x := math.Sqrt(float64(f))
			fmt.Println(x)
		}(i, &wg)
	}
	wg.Wait()
}
