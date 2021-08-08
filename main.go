package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Code  string
	Name  string
	Price float32
}

type Basket struct {
	Code             uint
	ProductsInBasket []Product
	TotalPrice       float32
}

var Products []Product
var Baskets []Basket

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Lana Merchandising</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func newBasket(w http.ResponseWriter, r *http.Request) {
	var basketIndex uint = 0

	if len(Baskets) > 0 {
		basketIndex := Baskets[len(Baskets)-1].Code
		basketIndex++
	}

	Baskets = append(Baskets, Basket{basketIndex, nil, 0.0})
}

func addProductToBasket(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func getTotalAmountInBasket(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func deleteBasket(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func handleRequests() {
	gorillaRouter := mux.NewRouter().StrictSlash(true)
	gorillaRouter.HandleFunc("/", index)
	gorillaRouter.HandleFunc("/about", about)
	log.Fatal(http.ListenAndServe(":3000", gorillaRouter))
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
