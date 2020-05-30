package main

import (
	"fmt"
	"math/rand"
)

//close channel , values returned in channel output
//range channel
//len/cap channel
//channels directions chan <-  // <- chan
// channel zero value
func main() {
	//closeChannelExample()
	//rangeChannelExample()
	//channelDirectionExample()
	//sender := make(chan string)
	var sender chan string
	fmt.Printf("%v", sender)
}

func channelDirectionExample() {
	sender := make(chan string, 1)
	receiver := make(chan string, 1)
	send(sender, "hello in the gateway")
	receiveAndSend(sender, receiver)
	msg := <-receiver

	fmt.Println(msg)
}

func send(sender chan<- string, msg string) {
	sender <- msg
}

func receiveAndSend(sender <-chan string, receiver chan<- string) {
	msg := <-sender
	receiver <- msg
}

func rangeChannelExample() {
	c := make(chan int, 3)
	rangeChannels(c)

	for _ = range c {
		//fmt.Print(val)
		//fmt.Print(" ")
		fmt.Printf("len %d, cap %d", len(c), cap(c))
		fmt.Println()
	}

	// this will return zero since range read all the values inside the channel
	// every read action in channel will decrement the len in 1.
	// The capacity is always the same.
	i := <-c
	fmt.Println(i)
	f := <-c
	fmt.Println(f)
}

func rangeChannels(c chan int) {
	c <- 1
	fmt.Printf("len %d, cap %d", len(c), cap(c))
	fmt.Println()
	c <- 2
	fmt.Printf("len %d, cap %d", len(c), cap(c))
	fmt.Println()
	c <- 3
	fmt.Printf("len %d, cap %d", len(c), cap(c))
	fmt.Println()
	close(c)
}

func closeChannelExample() {
	randChannel := make(chan int)
	go rands(randChannel)
	var count int
	for {
		i, ok := <-randChannel
		count++
		if !ok {
			break
		}

		fmt.Printf("index %d, value: %d \n", count, i)
	}
	fmt.Println("out of the loop")
}

func rands(c chan int) {
	for i := 0; i < 10; i++ {
		c <- rand.Int()
	}
	close(c)
}
