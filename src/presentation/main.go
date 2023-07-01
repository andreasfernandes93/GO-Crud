package main

import (
	"go-crud/src/presentation/controller"
	"net/http"

	
)

func main() {

	//rota index e controler
	http.HandleFunc("/products", controller.GetProducts)

	// criando servidor local na porta 8000
	http.ListenAndServe(":8000", nil)
}
