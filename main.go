package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

//Structs

type Product struct {
	gorm.Model

	Code  string
	Name  string
	Price float32
}

type Basket struct {
	gorm.Model

	ProductsInBasket []Product
}

type Total struct {
	Items []string
	Total float32
}

//Global variables

var Baskets []Basket
var db *gorm.DB
var dberr error

var (
	Products = []Product{
		{Code: "PEN", Name: "Lana Pen", Price: 5.0},
		{Code: "TSHIRT", Name: "Lana T-Shirt", Price: 20.0},
		{Code: "MUG", Name: "Lana Coffee Mug", Price: 7.5},
	}
)

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

//Products

func getAllProducts(w http.ResponseWriter, r *http.Request) {

}

func getProduct(w http.ResponseWriter, r *http.Request) { //Product {

}

//Creates a new basket with no products and a unique identifier
func newBasket(w http.ResponseWriter, r *http.Request) {

}

//Finds a basket by id and adds a product to it
func addProductToBasket(w http.ResponseWriter, r *http.Request) {
	/* 	vars := mux.Vars(r)
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
	   	} */

}

//Calculates basket total by adding products and applying discounts
func getTotalAmountInBasket(w http.ResponseWriter, r *http.Request) {
	/*	vars := mux.Vars(r)
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
	*/
}

//Finds a basket by id and deletes it
func deleteBasket(w http.ResponseWriter, r *http.Request) {

}

/////
//Init functions
/////

func handleRequests() {
	gorRouter := mux.NewRouter()
	db, dberr = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=postgres") //TODO: add password security

	if dberr != nil {
		fmt.Println(dberr)
		panic("Failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Basket{})

	for index := range Products {
		db.Create(&Products[index])
	}

	for index := range Baskets {
		db.Create(&Baskets[index])
	}

	gorRouter.HandleFunc("/", index)
	gorRouter.HandleFunc("/about", about)

	gorRouter.HandleFunc("/Products", about).Methods("GET")
	gorRouter.HandleFunc("/Products/{id}", about).Methods("GET")

	gorRouter.HandleFunc("/Baskets", newBasket).Methods("POST")
	gorRouter.HandleFunc("/Baskets/{id}", addProductToBasket).Methods("POST")
	gorRouter.HandleFunc("/Baskets/{id}", getTotalAmountInBasket)
	gorRouter.HandleFunc("/Baskets/{id}", deleteBasket).Methods("DELETE")

	handler := cors.Default().Handler(gorRouter)

	log.Fatal(http.ListenAndServe(":3000", handler))
}

//Applies discounts
/*func calculateTotal(id int) Total {

	var total Total

	for _, n := range Baskets[id].ProductsInBasket {
		total.Items = append(total.Items, n.Name)
		total.Items = append(total.Items, n.Name)
	}

	return total
}*/

////
//Main
////

func main() {

	fmt.Println("Server Starting...")

	fmt.Println("Loading Products...")

	fmt.Println("Starting Router...")
	handleRequests()
}
