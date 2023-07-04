package main

import (
	"log"
	"net/http"

	"github.com/andreasfernandes93/go-crud/src/presentation/routes"
)

func main() {

	fs := http.FileServer(http.Dir("scripts"))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", fs))
	//rota index e controler
	routes.Routes()

	// criando servidor local na porta 8000
	port := ":8000"
	log.Printf("Servidor iniciado. Conexão estabelecida na porta %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
	
}
