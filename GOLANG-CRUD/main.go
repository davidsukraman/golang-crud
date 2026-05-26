package main

import (
	"golang-crud/config"
	"golang-crud/controllers/categorycontroller"
	"golang-crud/controllers/homecontroller"
	"golang-crud/controllers/produkcontroller"
	"log"
	"net/http"
)

func main() {

	config.ConnectDB()

	//1. Home Page
	http.HandleFunc("/", homecontroller.Welcome)

	//2. Categories Page
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	//3. Products Page
	http.HandleFunc("/produk", produkcontroller.Index)
	http.HandleFunc("/produk/add", produkcontroller.Add)
	http.HandleFunc("/produk/edit", produkcontroller.Edit)
	http.HandleFunc("/produk/delete", produkcontroller.Delete)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)

}
