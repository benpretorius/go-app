package main

import (
	"fmt"
	"net/http"

	"github.com/benpretorius/go-app/pkg/handlers"
)

const portNumber = ":8080"

// main is the entry point for our application
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil) // The first 1024 ports is privileged. In other words you need to be a super user.

}
