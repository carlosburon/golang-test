package main

import (
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Code  string
	Name  string
	Price float32
}

var Products []Product

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

	fmt.Println("Loading Products...")
	Products = []Product{
		{Code: "PEN", Name: "Lana Pen", Price: 5.0},
		{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.0},
		{Code: "MUG", Name: "Lana Coffee Mug", Price: 7.5},
	}

	fmt.Println("Starting Router...")
	handleRequests()
}
