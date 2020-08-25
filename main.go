package main

import (
	"log"
	"fmt"
)

var bindAddress = env.string("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {
	
	env.Parse()

	l := log.New(os.Stdout, "products-api", log.LstdFlags)

	// create the handlers
	ph := handler.NewProducts(l)
}