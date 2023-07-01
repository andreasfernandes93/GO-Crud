package repository

import (
	"go-crud/src/domain/entity"
	database "go-crud/src/infra/db"
)

func GetAllProducts() ([]*entity.Product, error) {
    db := database.DbConection()
	defer db.Close()

    query := "SELECT * FROM Products"
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    products := []*entity.Product{}
    for rows.Next() {
        var id, quantity int
        var name, description string
        var price float64

        err := rows.Scan(&id, &name, &description, &price, &quantity)
        if err != nil {
            return nil, err
        }

        p := entity.NewProduct(id, name, description, price, quantity)
        products = append(products, p)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return products, nil
}
