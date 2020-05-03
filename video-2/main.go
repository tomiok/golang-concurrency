package main

import (
	"fmt"
	"runtime"
	"time"
)

//goroutines advantages

//They are lightweight.
//Ability to scale seamlessly.
//They are virtual threads. (One OS thread can multiplex many goroutines)
//Less memory requirement (2KB).
// Goroutines can communicate each other using channels

// Sequential execution = 7.093071521s

// Concurrent execution = 5.004216349s

func main() {


	numGR := runtime.NumGoroutine()
	fmt.Println(numGR)

	now := time.Now()

	//1st goroutine
	go func() {
		for i := 10; i < 20; i++ {
			time.Sleep(100 * time.Millisecond)
			//fmt.Println(i)
		}
		numGRf := runtime.NumGoroutine()
		fmt.Println(numGRf)
	}()

	//2nd goroutine
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
			//fmt.Println(i)
		}

		numGRs := runtime.NumGoroutine()
		fmt.Println(numGRs)
	}()

	//go longAndHeavyTask()

	elapsed := time.Since(now)

	fmt.Println(elapsed.String())
	time.Sleep(3 * time.Second)
}

func longAndHeavyTask() {
	time.Sleep(2 * time.Second)
}

func main2() {
	now := time.Now()

	for i := 10; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}

	longAndHeavyTask()

	elapsed := time.Since(now)

	fmt.Println(elapsed.String())
}
