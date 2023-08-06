package main

import (
	devTools "github.com/andreasfernandes93/go-crud/src/devUtils"
	"github.com/andreasfernandes93/go-crud/src/presentation/routes"
)

func main() {

	routes.Routes()

	devTools.StartServer(":9000")

}
