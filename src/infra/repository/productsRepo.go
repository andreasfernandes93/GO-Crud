package repository

import (
	"github.com/andreasfernandes93/go-crud/src/domain/entity"
	database "github.com/andreasfernandes93/go-crud/src/infra/db"
)

func GetAllProducts() ([]*entity.Product, error) {
    db := database.DbConection()
    defer db.Close()

    query := "SELECT * FROM Products"
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }

    products := []*entity.Product{}
    for rows.Next() {
        var id, quantity int
        var name, description string
        var price float64

        err := rows.Scan(&id, &name, &description, &price, &quantity)
        if err != nil {
            return nil, err
        }

        p := &entity.Product{
            Id:          id,
            Name:        name,
            Description: description,
            Price:       price,
            Quantity:    quantity,
        }
        products = append(products, p)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return products, nil
}

/*
func GetProduct(id string) (*entity.Product, error) {
	db := database.DbConection()
	defer db.Close()
	queryProductId := "SELECT * FROM Products WHERE id = ?"
	row := db.QueryRow(queryProductId, id)
}
*/

func InsertProduct(name string, description string, price float64, quantity int) {
	db := database.DbConection()
	defer db.Close()

	insertDb, err := db.Prepare("INSERT INTO Products (name, description, price, quantity) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	insertDb.Exec(name, description, price, quantity)
}

func DeleteProduct(id string) {
	db := database.DbConection()
	defer db.Close()

	deleteDb, err := db.Prepare("DELETE FROM Products WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	deleteDb.Exec(id)
}
