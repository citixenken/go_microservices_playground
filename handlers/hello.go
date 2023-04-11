package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// signature
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	h.l.Println("Hello, gopher!")
	data, err := io.ReadAll(r.Body)

	if err != nil {
	// rw.WriteHeader(http.StatusBadRequest)
	// rw.Write([]byte("Oops! Something went wrong"))
	http.Error(rw, "Oops! Something went wrong", http.StatusBadRequest)
	return
	}

	// log.Printf("Data: %s\n", data)
	fmt.Fprintf(rw, "Hello, %s", data)
}