package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	calls = []string{}
	stats = map[string]int{}
)

func main() {
	http.HandleFunc("/hello", greetUser)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func greetUser(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	calls = append(calls, name)

	nameStats, ok := stats[name]
	if !ok {
		stats[name] = 1
	}
	nameStats++
	stats[name] = nameStats

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "Hello, ", name)
	fmt.Printf("calls: %#v\n", calls)
	fmt.Printf("stats: %#v\n\n", stats)
}
