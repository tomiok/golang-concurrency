package main

import "fmt"

func main() {
	sender := make(chan string, 1)
	receiver := make(chan string, 1)

	send(sender, "hola")

	receiveAndSend(sender, receiver)
	msg := <-receiver

	fmt.Println(msg)
}

func send(c chan<- string, msg string) {
	c <- msg
}

func receiveAndSend(r <-chan string, s chan<- string) {
	msg := <-r
	s <- msg
}
