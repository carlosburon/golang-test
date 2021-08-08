package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Code  string  `json:"Code"`
	Name  string  `json:"Name"`
	Price float32 `json:"Price"`
}

type Basket struct {
	Id               uint `json:"Id"`
	ProductsInBasket []Product
}

var Products []Product
var Baskets []Basket

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Lana Merchandising (Homepage Endpoint)</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>SRE-Challenge: attempt by Carlos Buron</h1>")
}

func newBasket(w http.ResponseWriter, r *http.Request) {
	var basketIndex uint = 0
	var newProducts []Product
	var newBasket Basket

	if len(Baskets) > 0 {
		basketIndex = Baskets[len(Baskets)-1].Id
		basketIndex = basketIndex + 1
	}

	newBasket = Basket{basketIndex, newProducts}
	Baskets = append(Baskets, newBasket)

	json.NewEncoder(w).Encode(newBasket)

}

func addProductToBasket(w http.ResponseWriter, r *http.Request) {

}

func getTotalAmountInBasket(w http.ResponseWriter, r *http.Request) {

}

func deleteBasket(w http.ResponseWriter, r *http.Request) {

}

func handleRequests() {
	gorillaRouter := mux.NewRouter().StrictSlash(true)
	gorillaRouter.HandleFunc("/", index)
	gorillaRouter.HandleFunc("/about", about)
	gorillaRouter.HandleFunc("/basket", newBasket).Methods("POST")
	gorillaRouter.HandleFunc("/basket/{id}", addProductToBasket).Methods("POST")
	gorillaRouter.HandleFunc("/basket/{id}", getTotalAmountInBasket)
	gorillaRouter.HandleFunc("/basket/{id}", deleteBasket).Methods("DELETE")
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
