package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Structs

type Product struct {
	Code  string `json:"Code"`
	Name  string `json:"Name"`
	Price string `json:"Price"`
}

type Basket struct {
	Id               int `json:"Id"`
	ProductsInBasket []Product
}

type Total struct {
	Items []string
	Total float32
}

//Global variables

var Products []Product
var Baskets []Basket

/////
//Default handlers
/////

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Lana Merchandising (Homepage Endpoint)</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>SRE-Challenge Lana attempt: Carlos Buron</h1>")
}

/////
//Rest API handlers
/////

//Creates a new basket with no products and a unique identifier
func newBasket(w http.ResponseWriter, r *http.Request) {
	var basketIndex int = 0
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

//Finds a basket by id and adds a product to it
func addProductToBasket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	i, err := strconv.Atoi(key)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newProduct Product

	json.Unmarshal(reqBody, &newProduct)

	if err == nil {
		resultBasketId := searchBasket(i)

		if resultBasketId != len(Baskets) {
			Baskets[resultBasketId].ProductsInBasket = append(Baskets[resultBasketId].ProductsInBasket, newProduct)
			json.NewEncoder(w).Encode(Baskets[resultBasketId])
		} else {
			//TODO: handle basket id not found error
		}
	} else {
		//TODO: handle malformed basket id error
	}

}

//Calculates basket total by adding products and applying discounts
func getTotalAmountInBasket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	i, err := strconv.Atoi(key)

	if err == nil {
		resultBasketId := searchBasket(i)

		if resultBasketId != len(Baskets) {

			Baskets[resultBasketId].ProductsInBasket = append(Baskets[resultBasketId].ProductsInBasket, newProduct)
			json.NewEncoder(w).Encode(Baskets[resultBasketId])

		} else {
			//TODO: handle basket id not found error
		}
	} else {
		//TODO: handle malformed basket id error
	}

}

//Finds a basket by id and deletes it
func deleteBasket(w http.ResponseWriter, r *http.Request) {

}

//Request function for clarity in main

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

//////
//Basket auxiliary functions
//////

//Searchs a basket by id in the global array
func searchBasket(id int) int {
	for i, n := range Baskets {
		if id == n.Id {
			return i
		}
	}
	return len(Baskets)
}

//Applies discounts
func calculateTotal(id int) Total {

	var total Total

	for _, n := range Baskets[id].ProductsInBasket {
		total.Items = append(total.Items, n.Name)
		total.Items = append(total.Items, n.Name)
	}

	return total
}

////
//Main
////

func main() {

	fmt.Println("Server Starting...")

	fmt.Println("Loading Products...")
	Products = []Product{
		{Code: "PEN", Name: "Lana Pen", Price: "5.0"},
		{Code: "TSHIRT", Name: "Lana T-Shirt", Price: "20.0"},
		{Code: "MUG", Name: "Lana Coffee Mug", Price: "7.5"},
	}

	fmt.Println("Starting Router...")
	handleRequests()
}
