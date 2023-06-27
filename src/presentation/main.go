package main

import (
	"database/sql"
	"go-crud/src/domain/entity"
	"log"
	"net/http"
	"os"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

func dbConection() *sql.DB {

	// caminho do arquivo de banco de dados
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//conecção com o banco sqlite
	connection := currentDir + "/src/infra/db/db.db"
	db, err := sql.Open("sqlite3", connection)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

var pages = template.Must(template.ParseGlob("template/*.html"))

func main() {

	//rota index
	http.HandleFunc("/", index)

	// criando servidor local na porta 8000
	http.ListenAndServe(":8000", nil)
}

// manipulador
func index(w http.ResponseWriter, r *http.Request) {
	
	//abrindo conecção com db
	db := dbConection()

	//criando uma query
	selectProducts, err := db.Query("SELECT * FROM Products")
	if err != nil {
		panic(err.Error())
	}

	p := entity.Product{}
	products := []entity.Product{}

	//Scaneando o conteudo da tabela
	for selectProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	pages.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}
