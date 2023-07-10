package routes

import (
	"net/http"

	"github.com/andreasfernandes93/go-crud/src/presentation/controller"
)

func Routes() {

	http.HandleFunc("/", controller.GetProducts)
	http.HandleFunc("/add", controller.Add)
	http.HandleFunc("/insert", controller.CreateProduct)
	http.HandleFunc("/delete", controller.DeleteProduct)
	http.HandleFunc("/edit", controller.Edit)
	http.HandleFunc("/update", controller.UpdateProduct)

}
