package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		j := 0
		for ; j < 10; j++ {
			ping(j)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(5 * time.Second)

	fmt.Println("now show numbers")

	i := 0
	for ; i < 10; i++ {
		go func(number int) {
			fmt.Println(fmt.Sprintf("number: %d", number))
		}(i)
	}
	time.Sleep(2 * time.Second)
}

func ping(i int) {
	fmt.Println(fmt.Sprintf("sending ping: %d", i))
}
