package main

import (
	"example/go_microservices_playground/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// greedy pattern matching
	// http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request){
	// 	log.Println("Goodbye, gopher!")
	// })
	http.ListenAndServe(":9090", sm)
}