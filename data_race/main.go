package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	const gr = 100

	var wg sync.WaitGroup

	wg.Add(gr * 2)

	var n int = 0

	//Declaring mutex
	var m sync.Mutex
	for i := 0; i < gr; i++ {
		fmt.Println("No. of Goroutines after entering for loop:", runtime.NumGoroutine())
		go func() {
			time.Sleep(time.Second / 10)
			fmt.Println("Goroutines before lock (increment):", runtime.NumGoroutine())

			//Lock - only 1 will access
			m.Lock()
			n++
			fmt.Println("Goroutines in action (increment):", runtime.NumGoroutine())
			//Unlocking
			m.Unlock()
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			fmt.Println("Goroutines before lock (decrement):", runtime.NumGoroutine())
			m.Lock()
			n--
			fmt.Println("Goroutines in action (decrement):", runtime.NumGoroutine())
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(n)

}
