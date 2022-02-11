// HTTP server example
package main

import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler returns a greeting
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Good evening All from Aditya !!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
