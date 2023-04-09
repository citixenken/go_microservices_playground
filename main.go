package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		log.Println("Hello, gopher!")
		data, err := io.ReadAll(r.Body)

		if err != nil {
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Oops! Something went wrong"))
			http.Error(rw, "Oops! Something went wrong", http.StatusBadRequest)
			return
		}

		// log.Printf("Data: %s\n", data)
		fmt.Fprintf(rw, "Hello, %s", data)
	})
	// greedy pattern matching
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request){
		log.Println("Goodbye, gopher!")
	})
	http.ListenAndServe(":9090", nil)
}