package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello from Go!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
