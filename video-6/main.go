package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// wait group
func main() {
	var (
		wg       sync.WaitGroup
		budgets  []Budget
		ors      []OfficeRevenue
		users    []User
		response Response
	)
	done := make(chan bool)
	errChan := make(chan error)

	wg.Add(1) //this number is the same as the number of goroutines to be executed.
	go func() {
		budgets = getBudgets()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		officeRevenues, err := getOfficeRevenues()
		ors = officeRevenues

		if err != nil {
			errChan <- err
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		users = getUsers()
		wg.Done()
	}()

	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		response = Response{
			Users:         users,
			OfficeRevenue: ors,
			Budget:        budgets,
		}
		close(done)
	case err := <-errChan:
		fmt.Println(err.Error())
		close(errChan)
		return
	}

	fmt.Println(response)
}

type Response struct {
	Users         []User
	OfficeRevenue []OfficeRevenue
	Budget        []Budget
}

func getBudgets() []Budget {
	return []Budget{{
		City:            "Rosario",
		EstimatedBudget: 1500000,
		Variation:       10,
	}, {
		City:            "Rio de Janeiro",
		EstimatedBudget: 25000,
		Variation:       35,
	},
	}
}

func getOfficeRevenues() ([]OfficeRevenue, error) {
	if time.Now().Unix()%2 == 0 {
		return nil, errors.New("some error code")
	}

	return []OfficeRevenue{{City: "Rosario", Revenue: 1000000},
		{City: "Buenos Aires", Revenue: 200000}, {City: "Medellin", Revenue: 100000}}, nil
}

func getUsers() []User {
	return []User{{ID: 1, Name: "James"}, {ID: 2, Name: "Sue"}, {ID: 3, Name: "Bob"}}
}

type User struct {
	ID   int
	Name string
}

type OfficeRevenue struct {
	City    string
	Revenue int
}

type Budget struct {
	City            string
	EstimatedBudget int
	Variation       int
}
