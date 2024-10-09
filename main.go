package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, welcome to the Go web server!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server at port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
