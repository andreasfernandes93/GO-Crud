package controller

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/andreasfernandes93/go-crud/src/infra/repository"
)

var pages = template.Must(template.ParseGlob("template/*.html"))

func GetProducts(w http.ResponseWriter, r *http.Request) {

	selectProducts, err := repository.GetAllProducts()
	if err != nil {
		log.Fatal(err)
	}

	pages.ExecuteTemplate(w, "Index", selectProducts)
}

func Add(w http.ResponseWriter, r *http.Request) {
	pages.ExecuteTemplate(w, "Add", nil)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Fatal("Erro na conversão do preço", err)
		}

		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Fatal("Erro na conversão da quantidade", err)
		}

		repository.InsertProduct(name, description, priceFloat, quantityInt)
	}
	http.Redirect(w, r, "/", 301)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	repository.DeleteProduct(idProduct)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := repository.GetProduct(idProduct)
	pages.ExecuteTemplate(w, "Edit", product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal("Erro na conversão do ID", err)
		}

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Fatal("Erro na conversão do preço", err)
		}

		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Fatal("Erro na conversão da quantidade", err)
		}

		repository.UpdateProduct(idInt, name, description, priceFloat, quantityInt)
	}

	http.Redirect(w, r, "/", 301)
}
