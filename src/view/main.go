package main

import (
	"go-crud/src/model/entity"
	"net/http"
	"text/template"
)

var pages = template.Must(template.ParseGlob("template/*.html"))

func main() {
	//rota index
	http.HandleFunc("/", index)

	// criando servidor local na porta 8000
	http.ListenAndServe(":8000", nil)
}

// manipulador
func index(w http.ResponseWriter,r *http.Request){
	products := []entity.Product{
		{Name: "Notebook", Description: "CoreI3 | 8GB RAM | SSD 1TB", Price: 2000.00, Quantity: 200},
		{Name: "Smartphone", Description: "6.5'' Display | 128GB Storage | 5000mAh Battery", Price: 1200.00, Quantity: 150},
		{Name: "TV", Description: "55'' 4K Ultra HD | Smart TV | HDR", Price: 2500.00, Quantity: 100},
		{Name: "Headphones", Description: "Wireless | Noise Cancelling | Bluetooth", Price: 150.00, Quantity: 300},
	}	

	pages.ExecuteTemplate(w, "Index", products)
}