package controller

import (
	"go-crud/src/infra/repository"
	"net/http"
	"text/template"
)

var pages = template.Must(template.ParseGlob("template/*.html"))

// GetProducts retrieves all products from the database and renders them on the index page.
//
// Parameters:
// - w: the http.ResponseWriter used to write the response.
// - r: the *http.Request representing the incoming request.
//
// Return type:
// None.
func GetProducts(w http.ResponseWriter, r *http.Request) {

	//criando uma query
	selectProducts, err := repository.GetAllProducts()
	if err != nil {
		panic(err.Error())
	}

	pages.ExecuteTemplate(w, "Index", selectProducts)
}
