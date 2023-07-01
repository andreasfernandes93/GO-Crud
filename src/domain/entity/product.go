package entity

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func NewProduct(id int, name string, description string, price float64, quantity int) *Product {
	return &Product{
		Id:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
}
