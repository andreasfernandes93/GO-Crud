package main

import (
	"net/http"

	"github.com/andreasfernandes93/go-crud/src/presentation/routes"
)

func main() {

	//rota index e controler
	routes.Routes()

	// criando servidor local na porta 8000
	http.ListenAndServe(":8000", nil)
}
