package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page.")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2. %d", sum))
}

// addValues adds to integers and returns the sum
func addValues(x, y int) int {

	return x + y
}

// main entry point for our application
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil) // The first 1024 ports is privileged. In other words you need to be a super user.

}
