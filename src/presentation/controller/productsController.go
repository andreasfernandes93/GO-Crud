package controller

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/andreasfernandes93/go-crud/src/infra/repository"
)

var pages = template.Must(template.ParseGlob("template/*.html"))

// GetProducts retrieves all products and executes the "Index" template.
//
// It takes in two parameters, w of type http.ResponseWriter and r of type *http.Request.
// It does not return any value.
func GetProducts(w http.ResponseWriter, r *http.Request) {

	selectProducts, err := repository.GetAllProducts()
	if err != nil {
		panic(err)
	}

	pages.ExecuteTemplate(w, "Index", selectProducts)
}

func CreateFormProduct(w http.ResponseWriter, r *http.Request) {
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
