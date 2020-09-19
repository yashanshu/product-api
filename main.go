package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/yashanshu/product-api/handlers"
)

func main() {

	l := log.New(os.Stdout, "products-api", log.LstdFlags)

	// create the handlers
	ph := handlers.NewProducts(l)

	// create a new serve mux and register the handlers.
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// create a new server
	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	err := s.ListenAndServe()
	if err != nil {
		l.Printf("Error starting the server: %s\n", err)
	}
}
