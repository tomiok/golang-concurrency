package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

//1. st output the questions
//2. st input the answers
//1 & 2 in a goroutine
//3 check the answer
//4 a select statement with a timer and a flag when is done
//5 print the results
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	questions := getQuestions()
	counter := 0
	done := make(chan bool)
	go func() {
		for _, question := range questions {
			fmt.Print(question[0], ":")
			_ = scanner.Scan()
			answer := scanner.Text()
			if answer == question[1] {
				counter++
			}
		}
		done <- true
	}()

	gameEngine(done, time.NewTicker(20*time.Second))
	fmt.Printf("Your score is: %d", counter)
}

func gameEngine(done chan bool, ticker *time.Ticker) {
	select {
	case <-done:
		fmt.Println()
		fmt.Println("Game finished")
	case <-ticker.C:
		fmt.Println()
		fmt.Println("outta of time dude! you lost!")
	}
}

func getQuestions() [][]string {
	return [][]string{{"10+10", "20"}, {"5+3", "8"}, {"20+10", "30"}, {"10+15", "25"}}
}
