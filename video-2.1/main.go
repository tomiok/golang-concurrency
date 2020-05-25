package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	s := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/async", func(w http.ResponseWriter, r *http.Request) {
		go longTask("tomi")
		t := time.Now().String()
		_, _ = w.Write([]byte(fmt.Sprintf("the time is %s", t)))
	})

	log.Println("http working fine in port 8080")
	log.Fatal(s.ListenAndServe())
}

func longTask(s string) {
	time.Sleep(5 * time.Second)
	log.Println("important log from: " + s)
}
