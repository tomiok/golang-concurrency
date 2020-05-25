package main

import (
	"fmt"
)

func main() {
	fmt.Println("calling worker")
	done := make(chan bool)

	go worker(done)

	b := <-done

	fmt.Println(b)
}

func worker(done chan bool) {
	fmt.Println("inside de worker")
	done <- true
}

func bufferedChannel() {
	c := make(chan string, 2)

	c <- "one"
	c <- "two"

	fmt.Println(<-c)
	fmt.Println(<-c)
}

func firstExample() {
	c := make(chan string)

	go func() { c <- "hello from channel" }()

	res := <-c

	fmt.Println(res)
}
