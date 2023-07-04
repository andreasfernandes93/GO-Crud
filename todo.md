
(A) Implements DTO of Product entity @Entity
(A) Implements VOs in Product entity @Entity
(A) Implements Controllers in @Presentation
(B) Correct the routes files @Presentation


// rascunho
func GetProduct(id string) *entity.Product {
	db := database.DbConection()
	defer db.Close()
	queryProductId := "SELECT * FROM Products WHERE id = ?"
    row := db.QueryRow(queryProductId, id)
    product := &entity.Product{}
    err := row.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity)
    if err != nil {
        return nil
    }
 
    return product
}