package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Version2. Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting hello server...")
	http.ListenAndServe(":8080", nil)
}
