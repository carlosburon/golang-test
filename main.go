package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

//Structs

type Product struct {
	gorm.Model

	Code  string `gorm:"unique"`
	Name  string
	Price float32
}

type ProductRequest struct {
	Code     string
	Quantity string
}

type Basket struct {
	gorm.Model

	ProductsInBasket string
}

type Total struct {
	Items string
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
	var allProducts []Product

	db.Find(&allProducts)

	json.NewEncoder(w).Encode(&allProducts)
}

func getProduct(w http.ResponseWriter, r *http.Request) { //Product {
	vars := mux.Vars(r)
	key := vars["id"]
	var product Product

	db.First(&product, key)

	json.NewEncoder(w).Encode(&product)

}

//Creates a new basket with no products and a unique identifier
func newBasket(w http.ResponseWriter, r *http.Request) {
	var result string
	newBasket := Basket{gorm.Model{}, ""}
	dbc := db.Create(&newBasket)

	if dbc.Error != nil {
		log.Panic("Error creating basket!")
		result = "Error"
	} else {
		result = "Basket created successfully"
	}

	json.NewEncoder(w).Encode(result)

}

//Returns all baskets in the database
func getAllBaskets(w http.ResponseWriter, r *http.Request) {
	var allBaskets []Basket

	db.Find(&allBaskets)

	json.NewEncoder(w).Encode(&allBaskets)
}

//Finds a basket by id and adds a product to it
func addProductToBasket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var basket Basket
	var product Product
	var productRequest ProductRequest
	var expandedBasket []string
	var quantityCasted int
	var errCasting error

	//Finds the product in the database by code
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &productRequest)
	db.Where("Code = ?", productRequest.Code).First(&product)

	//Finds the basket
	db.First(&basket, key)

	//Adds the product to the basket
	if basket.ProductsInBasket != "" {
		expandedBasket = strings.Split(basket.ProductsInBasket, ",")
	}

	quantityCasted, errCasting = strconv.Atoi(productRequest.Quantity)
	if errCasting != nil {
		quantityCasted = 0
		log.Panic("Cannot convert quantity of products added. Set to 0")
	}

	if product.Code != "" {
		for i := 1; i <= quantityCasted; i++ {
			expandedBasket = append(expandedBasket, product.Code)
		}
		basket.ProductsInBasket = strings.Join(expandedBasket[:], ",")
		db.Save(&basket)
	} else {
		log.Panic("Invalid product code. Not adding anything to basket")
	}

	//	json.NewEncoder(w).Encode(&product)
	json.NewEncoder(w).Encode(&basket)
}

//Calculates basket total by adding products and applying discounts
func getTotalAmountInBasket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var basket Basket
	var expandedBasket []string
	var total Total

	//TODO: do not hardcode discounts
	var penCounter int
	var tshirtCounter int

	//Finds the product in the database by code
	//db.Where("Code = ?", productRequest.Code).First(&product)

	//Finds the basket
	db.First(&basket, key)

	if basket.ProductsInBasket != "" {
		expandedBasket = strings.Split(basket.ProductsInBasket, ",")

		for index := range expandedBasket {
			var product Product
			db.Where("Code = ?", expandedBasket[index]).First(&product)

			//TODO: do not hardcode discounts
			if expandedBasket[index] == "PEN" {
				penCounter++
				if penCounter%2 == 0 {
					total.Total += 0
				} else {
					total.Total += product.Price
				}
			} else if expandedBasket[index] == "TSHIRT" {
				tshirtCounter++
				if tshirtCounter == 3 {
					total.Total -= product.Price * 0.25
					total.Total -= product.Price * 0.25
					total.Total += product.Price * 0.75
				} else if tshirtCounter > 3 {
					total.Total += product.Price * 0.75
				} else {
					total.Total += product.Price
				}
			} else {
				total.Total += product.Price
			}
		}
		total.Items = basket.ProductsInBasket
	} else {
		total.Items = "Basket empty"
		total.Total = 0
	}

	json.NewEncoder(w).Encode(&basket)
	json.NewEncoder(w).Encode(total)
}

//Finds a basket by id and deletes it
func deleteBasket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var basket Basket

	db.First(&basket, params["id"])
	json.NewEncoder(w).Encode(&basket)

	db.Delete(&basket)

}

/////
//Init functions
/////

func handleRequests() {
	gorRouter := mux.NewRouter()

	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
	)
	//Tries to reconnect to the database. Should implement it with os.Getenv("PGRETRIES") casted to int

	retries := 10

	for retries > 0 {

		db, dberr = gorm.Open("postgres", pgConnString)

		if dberr == nil {
			log.Println("Succesfully connected to database")
			break
		}
		log.Println(dberr)
		log.Printf("Failed to connect to database. %d tries left\n", retries)
		time.Sleep(5 * time.Second)
		retries--
	}

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

	gorRouter.HandleFunc("/Products", getAllProducts).Methods("GET")
	gorRouter.HandleFunc("/Products/{id}", getProduct).Methods("GET")

	gorRouter.HandleFunc("/Baskets", newBasket).Methods("POST")
	gorRouter.HandleFunc("/Baskets", getAllBaskets).Methods("GET")
	gorRouter.HandleFunc("/Baskets/{id}/items", addProductToBasket).Methods("POST")
	gorRouter.HandleFunc("/Baskets/{id}", getTotalAmountInBasket).Methods("GET")
	gorRouter.HandleFunc("/Baskets/{id}", deleteBasket).Methods("DELETE")

	handler := cors.Default().Handler(gorRouter)

	log.Fatal(http.ListenAndServe(":3000", handler))
}

////
//Main
////

func main() {

	log.Println("Server Starting...")

	log.Println("Loading Products...")

	log.Println("Starting Router...")
	handleRequests()
}
