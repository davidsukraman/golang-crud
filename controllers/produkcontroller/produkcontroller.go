package produkcontroller

import (
	"golang-crud/entities"
	"golang-crud/models/categorymodel"
	"golang-crud/models/produkmodel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {

	produk := produkmodel.GetAll()
	data := map[string]any{
		"produk": produk,
	}

	temp, err := template.ParseFiles("views/produk/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)

}

func Add(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/produk/create.html")
		if err != nil {
			panic(err)
		}

		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
		}
		temp.Execute(w, data)
	}

	if r.Method == "POST" {

		var produk entities.Produk

		categoryID, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			panic(err)
		}

		produk.Name = r.FormValue("name")
		produk.Category.Id = uint(categoryID)
		produk.Stok = int64(stock)
		produk.Keterangan = r.FormValue("description")
		produk.CreatedAt = time.Now()
		produk.UpdatedAt = time.Now()
		if ok := produkmodel.Create(produk); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/produk", http.StatusSeeOther)
	}

}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
