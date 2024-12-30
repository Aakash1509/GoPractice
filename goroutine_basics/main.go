package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 (goroutine) execution started")
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("f1,i=", i)
	}
	fmt.Println("f1 execution finished")
	wg.Done()
}

func f2() {
	fmt.Println("f2 (goroutine) execution started")
	for i := 5; i < 8; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("f2,i=", i)
	}
	fmt.Println("f2 execution finished")
}

func main() {
	//fmt.Println("Main execution started")
	//fmt.Println("No. of CPUs:", runtime.NumCPU())
	//fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
	//fmt.Println("OS:", runtime.GOOS)
	//fmt.Println("Arch:", runtime.GOARCH)
	//
	//fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	// 1.
	// Create a new instance of sync.WaitGroup (we’ll call it simply wg)
	// This WaitGroup is used to wait for all the goroutines that have been launched to finish.

	var wg sync.WaitGroup

	wg.Add(1)

	// 2.
	// Call wg.Add(n) method before attempting to
	// launch the go routine.

	wg.Add(1) //  n which is 1 is the number of goroutines to wait

	//Launching a goroutine
	go f1(&wg) //it takes in pointer to sync.WaitGroup//The main function actually terminates before go routine gets a chance to execute
	fmt.Println("No. of Goroutines after go f1():", runtime.NumGoroutine())
	f2()
	//without using wait group time.Sleep(time.Second * 2) //Need to introduce sleep
	//f1() gets chance to execute

	// 4.
	// Finally, we call wg.Wait()to block the execution of main() until the goroutines
	// in the WaitGroup have successfully completed.
	wg.Wait()

	fmt.Println("main execution finished")

	go func() {

	}()
}
