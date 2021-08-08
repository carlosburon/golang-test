package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Lana Merchandising</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func handleRequests() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {

	fmt.Println("Server Starting...")
	handleRequests()
}
