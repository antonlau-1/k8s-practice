package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func nginxHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://nginx")
	if err != nil {
		http.Error(w, "Failed to connect to NGINX server: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to retrieve response body"+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(body)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/nginx", nginxHandler)
	fmt.Println("Starting hello server...")
	http.ListenAndServe(":8080", nil)
}
