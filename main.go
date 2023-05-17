package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	http.HandleFunc("/api/demo", func(w http.ResponseWriter, r *http.Request) {
		person := Person{
			Name: "Anna",
			Age:  23,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(person)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
